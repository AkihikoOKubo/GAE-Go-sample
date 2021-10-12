package handler

import (
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/auth"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/bridge"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/output"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/user"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/log"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"
	// GAEではmain packageが無いのでここに置きます
	"github.com/alioygur/gores"
)

var (
	cnf     *config.ServerConfig
	uh      *user.Handler
	ah      *auth.Handler
	bh      *bridge.Handler
)

func init() {
	var err error

	if cnf, err = config.NewConfig(); err != nil {
		panic(fmt.Sprintf("error: %s", err.Error()))
	}

	uh = user.NewHandler(cnf)

	bh = bridge.NewHandler(cnf)

	ah = auth.NewHandler(cnf)
}

// Router creates NewRouter
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ping", ping).Methods(http.MethodGet)
	r.HandleFunc("/users", push(uh.Post)).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", push(uh.Get)).Methods(http.MethodGet)
	r.HandleFunc("/ping_other_service", push(bh.PingOtherService)).Methods(http.MethodGet) // 他のGAEサービスに向けてpingします
	// MEMO: 認証を試す場合はpush→mustAuthに変更してください
	return r
}

// fn はエンドポイントの処理のInterfaceを定義です
type fn func(context.Context, http.ResponseWriter, *http.Request)

// push はエンドポイントの処理のラッパーです
func push(fn fn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := loggingContext(r)
		fn(c, w, r)
	}
}

// mustAuth は要認証のEndpointのラッパーです。accessTokenを検証し、問題があれば401を返します
func mustAuth(fn fn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := loggingContext(r)
		err := ah.VerifyToken(c, r)
		if err != nil {
			e := model.NewWrapErr(err)
			gores.JSON(w, e.StatusCode, output.NewErrResponse(e))
			return
		}
		fn(c, w, r)
	}
}

// loggingContext はContextにセットアップ済みのLoggerをセットします。
// これにより、TraceIDやSpanIDがセットされ、またリクエストの開始時刻、処理秒数が設定された独自のアクセスログが"stderr"に出力されます。
// Cloud logging上で"stderr"を開くと、リクエスト単位でまとまって(折り畳まれて)かつSeverityが設定された状態のログを表示できます。
func loggingContext(r *http.Request) context.Context {
	ctx := log.SetupContextLogger(r)
	log.Logger(ctx).Info("Logger setup done.")
	return ctx
}

// Ping is ...
func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

