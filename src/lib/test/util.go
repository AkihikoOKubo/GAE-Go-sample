package test

import (
	"net/http"

	"github.com/gorilla/mux"
)

// MakeRouter テスト用
func MakeRouter(path, method string, fn func(http.ResponseWriter, *http.Request)) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc(path, fn).Methods(method)
	return r
}
