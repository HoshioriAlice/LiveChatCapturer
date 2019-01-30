package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logFile, err := os.OpenFile("livechat.log", os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println(err)
	} else {
		log.SetOutput(logFile)
	}
	err = http.ListenAndServe("0.0.0.0:"+os.Args[1], &Server{})
	log.Println(err)
}
