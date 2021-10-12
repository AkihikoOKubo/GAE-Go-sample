package model

import (
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"net/http"
)

// WrapErr はエラーオブジェクトをラップした、エラーレスポンス作成用の構造体です
type WrapErr struct {
	OrgError   error
	StatusCode int
	Message    string
}

// Error is ...
func (err *WrapErr) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.OrgError.Error(), err.StatusCode)
}

// NewWrapErr はWrapErrを作成します
func NewWrapErr(e error) WrapErr {
	var sc int
	// エラーの種別毎にエラーメッセージを設定します
	// Application独自のエラーを定義している場合や、他言語対応する場合などはここでメッセージを出し分けられます
	switch e {
	case config.ErrNotFound:
		sc = http.StatusNotFound
	case config.ErrUnauthorized:
		sc = http.StatusUnauthorized
	case config.ErrBadRequest:
		sc = http.StatusBadRequest
	default:
		sc = http.StatusInternalServerError
	}

	return WrapErr{
		OrgError:   e,
		StatusCode: sc,
		Message:    e.Error(),
	}
}
