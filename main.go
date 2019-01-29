package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	liveChatURL := "https://www.youtube.com/live_chat/get_live_chat?continuation=%s&hidden=false&pbj=1"
	ContinuationString := os.Args[1]
	for {
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
		dec := json.NewDecoder(resp.Body)
		var lc LiveChat
		dec.Decode(&lc)

		for _, action := range lc.Response.ContinuationContents.LiveChatContinuation.Actions {
			name := action.AddChatItemAction.Item.LiveChatTextMessageRenderer.AuthorName.SimpleText
			message := action.AddChatItemAction.Item.LiveChatTextMessageRenderer.Message.SimpleText
			fmt.Println(name+":", message)
		}
		if len(lc.Response.ContinuationContents.LiveChatContinuation.Continuations) > 0 {
			ContinuationString = lc.Response.ContinuationContents.LiveChatContinuation.Continuations[0].TimedContinuationData.Continuation
			//fmt.Println("ContinuationString Change To", ContinuationString)
		}
		time.Sleep(time.Second)
	}
}
