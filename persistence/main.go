package main

import (
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatalln("Argument count error")
	}
	file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("Cannot Open File")
	}
	defer file.Close()
	log.SetOutput(file)
	continuation := getFirstContinuationString(os.Args[1])
	if len(continuation) == 0 {
		log.Fatalln("Can not get livechat")
	}
	count := 0
	for {
		messages := getLiveChat(&continuation)
		for _, message := range messages {
			count++
			if len(message.Purchase) > 0 {
				log.Println(count, message.Sender, "["+message.Purchase+"]: ", message.Message)
			} else {
				log.Println(count, message.Sender+": ", message.Message)
			}
		}
		time.Sleep(time.Second)
	}
}
