package main

import (
	"net/http"
	"study/knowledge_planet/homework/20190430/handler"
)

func main() {
	http.HandleFunc("/index", handler.Index)
	http.ListenAndServe(":8080", nil)
}
