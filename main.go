package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080", // デプロイするなら環境変数から読むようにしたほうが良き？
	}
	http.HandleFunc("/pfc", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}
	v, ok := q["calory"]
	if !ok {
		// TODO: パラメーターの不足はクライアントエラーなのでlog出力してプロセスを終了するのではなく、400を返すようにする
		log.Fatal("calory param must be present")
	}
	fmt.Fprintln(w, v[0])
}
