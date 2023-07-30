package main

import (
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	log.Printf("请求测试")
	w.Write([]byte(string("HnadleFunc")))
}

func main() {
	http.HandleFunc("/sayHello", sayHello)
	http.ListenAndServe("localhost:8000", nil)
}
