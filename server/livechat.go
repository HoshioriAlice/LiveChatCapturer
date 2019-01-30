package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// LiveChatMessage is response of LiveChatUpdateRequest
type LiveChatMessage struct {
	Sender   string `json:"sender"`
	Message  string `json:"message"`
	Purchase string `json:"purchase"`
}

type LiveChatResponse struct {
	Continuation string            `json:"continuation"`
	Messages     []LiveChatMessage `json:"messages"`
}

type Server struct{}

func getFirstContinuationString(liveURL string) string {
	const pattern string = "{\"reloadContinuationData\":{\"continuation\":\""
	// 获取直播间网页
	client := &http.Client{}
	req, err := http.NewRequest("GET", liveURL, nil)
	if err != nil {
		log.Println("Request Build Failed", liveURL)
		return ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Host", "www.youtube.com")
	req.Header.Add("Accpet", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accpet-Encoding", "defalte")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Get Live Page Failed", liveURL)
		return ""
	}
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read Response Error", liveURL)
		return ""
	}
	// 找到Continuation信息
	begin := bytes.Index(s, []byte(pattern))
	if begin == -1 {
		log.Println("Search Continuation Information Error", liveURL)
		return ""
	}
	begin += len(pattern)
	end := bytes.Index(s[begin:], []byte("\""))
	if end == -1 {
		log.Println("Search Continuation Information Error", liveURL)
		return ""
	}
	return string(s[begin : begin+end])
}

func getLiveChat(continuation *string) []LiveChatMessage {
	liveChatURL := "https://www.youtube.com/live_chat/get_live_chat?continuation=%s&hidden=false&pbj=1"
	var result []LiveChatMessage
	URL := fmt.Sprintf(liveChatURL, *continuation)
	// 发送获取LiveChat的请求
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Println("Request Build Failed", URL)
		return result
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Host", "www.youtube.com")
	req.Header.Add("Accpet", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accpet-Encoding", "defalte")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Get LiveChat Failed")
		return result
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	// 解析json
	var lc LiveChat
	err = json.Unmarshal(buf, &lc)
	if err != nil {
		log.Println("Parse JSON Error, But Try To Continue", URL)
	}
	// 获取评论信息并打印
	for _, action := range lc.Response.ContinuationContents.LiveChatContinuation.Actions {
		if action.AddChatItemAction.Item.LiveChatTextMessageRenderer != nil {
			render := action.AddChatItemAction.Item.LiveChatTextMessageRenderer
			result = append(result, LiveChatMessage{Sender: render.AuthorName.SimpleText, Message: render.Message.SimpleText})
		} else if action.AddChatItemAction.Item.LiveChatPaidMessageRenderer != nil {
			render := action.AddChatItemAction.Item.LiveChatPaidMessageRenderer
			result = append(result, LiveChatMessage{Sender: render.AuthorName.SimpleText, Message: render.Message.SimpleText, Purchase: render.PurchaseAmountText.SimpleText})
		}
	}
	// 更新Continuation信息
	continuations := lc.Response.ContinuationContents.LiveChatContinuation.Continuations
	if len(continuations) > 0 {
		if continuations[0].TimedContinuationData != nil {
			*continuation = continuations[0].TimedContinuationData.Continuation
		} else if continuations[0].InvalidationContinuationData != nil {
			*continuation = continuations[0].InvalidationContinuationData.Continuation
		} else {
			log.Println("No Continuation Infomation", URL)
		}
	}
	return result
}

type ConnectResponse struct {
	Status       string `json:"status"`
	Continuation string `json:"continuation"`
}

func (c *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	action := r.Header.Get("Action")
	var buf []byte
	if action == "Connect" {
		livePageURL := r.Header.Get("Live-Page")
		if len(livePageURL) > 0 {
			continuation := getFirstContinuationString(livePageURL)
			if len(continuation) > 0 {
				log.Println("Get Live Page Success")
				buf, _ = json.Marshal(ConnectResponse{"Success", continuation})
			} else {
				log.Println("Get Live Page Failed", livePageURL)
				buf, _ = json.Marshal(ConnectResponse{"Failed", ""})
			}
		} else {
			log.Println("Get Live Page Error", livePageURL)
			buf, _ = json.Marshal(ConnectResponse{"URL Error", ""})
		}
	} else if action == "Update" {
		continuation := r.Header.Get("Continuation")
		if len(continuation) > 0 {
			messages := getLiveChat(&continuation)
			buf, _ = json.Marshal(LiveChatResponse{continuation, messages})
		} else {
			w.WriteHeader(404)
			return
		}
	} else {
		w.WriteHeader(404)
		return
	}
	_, err := w.Write(buf)
	if err != nil {
		log.Println("Send Response Error\n", buf)
	}
}
