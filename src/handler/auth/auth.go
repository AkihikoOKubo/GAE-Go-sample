package auth

import (
	"context"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"net/http"
)

// VerifyToken はヘッダに設定されたidTokenを検証し、結果を返却します。
// 以下処理はフロントでFirebase Authenticationを利用してidTokenとをセットしている場合のサンプルです。
// フロントに期待するHeader:
//     AppAuthorization : idTokenをセットしてください
//     ClientID : FirebaseIDをセットしてください
func (h *Handler) VerifyToken(ctx context.Context, r *http.Request) (err error) {
	// MEMO: IAPがHeader`Authorization`を使うので、アプリ独自のAuthenticationは別のヘッダを設定する必要があります
	// IAP側のHeader名を変更することもできるのですが未検証です
	t := r.Header.Get("AppAuthorization")
	if t == "" {
		return config.ErrUnauthorized
	}
	cid := r.Header.Get("ClientID")

	token := &model.Token{
		Token:     t,
		ClientID:  cid,
	}
	return h.AuthUsecase.VerifyToken(ctx, token)
}