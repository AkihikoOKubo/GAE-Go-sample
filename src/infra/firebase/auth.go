package firebase

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
)

// VerifyIDToken はidTokenを検証します
func (f *Firebase) VerifyIDToken(ctx context.Context, cli *auth.Client, t *model.Token) error {
	res, err := cli.VerifyIDToken(ctx, t.Token)
	if err != nil {
		fmt.Printf("err :%s", err.Error())
		return err
	}

	if t.ClientID != res.UID {
		return config.ErrUnauthorized
	}
	return nil
}
