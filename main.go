package main

import (
	"log"
	"net/http"
	"os"
	"shuhelperConfigStorage/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingPongHandler)
	http.HandleFunc("/config", handler.ConfigHandler)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
