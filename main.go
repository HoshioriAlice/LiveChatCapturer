package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	pattern string = "{\"reloadContinuationData\":{\"continuation\":\""
)

func getFirstContinuationString(liveURL string) string {
	// 获取直播间网页
	client := &http.Client{}
	req, err := http.NewRequest("GET", liveURL, nil)
	if err != nil {
		log.Fatalln("Request Build Failed")
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Add("Host", "www.youtube.com")
	req.Header.Add("Accpet", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Add("Accpet-Encoding", "defalte")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Get Live Page Failed")
	}
	s, err := ioutil.ReadAll(resp.Body)
	// 找到Continuation信息
	begin := bytes.Index(s, []byte(pattern)) + len(pattern)
	end := bytes.Index(s[begin:], []byte("\""))
	fmt.Println(begin, end)
	return string(s[begin : begin+end])
}

func main() {
	// 获取LiveChat的URL
	liveChatURL := "https://www.youtube.com/live_chat/get_live_chat?continuation=%s&hidden=false&pbj=1"
	ContinuationString := getFirstContinuationString(os.Args[1])
	for {
		// 发送获取LiveChat的请求
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf(liveChatURL, ContinuationString), nil)
		if err != nil {
			log.Fatalln("Request Build Failed")
		}
		req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
		req.Header.Add("Host", "www.youtube.com")
		req.Header.Add("Accpet", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		req.Header.Add("Accpet-Encoding", "defalte")
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Get LiveChat Failed")
		}
		// 解析json
		dec := json.NewDecoder(resp.Body)
		var lc LiveChat
		dec.Decode(&lc)
		// 获取评论信息并打印
		for _, action := range lc.Response.ContinuationContents.LiveChatContinuation.Actions {
			name := action.AddChatItemAction.Item.LiveChatTextMessageRenderer.AuthorName.SimpleText
			message := action.AddChatItemAction.Item.LiveChatTextMessageRenderer.Message.SimpleText
			fmt.Println(name+":", message)
		}
		// 更新Continuation信息
		if len(lc.Response.ContinuationContents.LiveChatContinuation.Continuations) > 0 {
			ContinuationString = lc.Response.ContinuationContents.LiveChatContinuation.Continuations[0].TimedContinuationData.Continuation
			//fmt.Println("ContinuationString Change To", ContinuationString)
		}
		time.Sleep(time.Second)
	}
}
