package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080", // デプロイするなら環境変数から読むようにしたほうが良き？
	}
	http.HandleFunc("/pfc", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hogehoge")
}
