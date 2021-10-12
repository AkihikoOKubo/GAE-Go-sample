package main

import (
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler"
	log "github.com/yfuruyama/stackdriver-request-context-log"
	"net/http"
	"os"
)

// main はアプリケーションの起動関数です。
func main() {
	// ハンドラとしてmux.Routerを仕込みます
	r := handler.Router()
	http.Handle("/", r)

	// Loggerをハンドラに仕込みます
	cnf, err := config.NewConfig()
	if err != nil {
		panic(fmt.Sprintf("error: %s", err.Error()))
	}
	config := log.NewConfig(cnf.App.ProjectID)
	handler := log.RequestLogging(config)(r)

	// GAE-2nd-Genからは以下のようにListen開始処理を書来ます
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Printf("Defaulting to port %s", port)
	}
	fmt.Printf("Listening on port %s", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Printf("err %s", err.Error())
	}
}