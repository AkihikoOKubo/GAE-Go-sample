package repository

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"golang.org/x/net/context"
)

// Firebase はFirebaseに関する実装を持ちます
type Firebase interface {
	NewFirestoreClint(ctx context.Context) (*firestore.Client, error)
	FinishFirestoreClient(cli *firestore.Client) error
	NewAuthClint(ctx context.Context) (cli *auth.Client, err error)
	VerifyIDToken(ctx context.Context, cli *auth.Client, t *model.Token) error
	User(ctx context.Context, cli *firestore.Client, id string) (*model.User, error)
	SaveUser(ctx context.Context, cli *firestore.Client, user *model.User) error
}