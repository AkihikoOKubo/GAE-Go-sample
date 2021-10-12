package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	fb "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
)

// Firebase はFirebaseに関する構造体です
type Firebase struct {
	Cnf *config.Firebase
}

func (f *Firebase) newApp(ctx context.Context) (*fb.App, error) {
	return fb.NewApp(ctx, &fb.Config{ProjectID: f.Cnf.ProjectID})
}

// NewFirestoreClint is return firestore client
func (f *Firebase) NewFirestoreClint(ctx context.Context) (*firestore.Client, error) {
	app, err := f.newApp(ctx)
	if err != nil {
		return nil, err
	}
	return app.Firestore(ctx)
}

// FinishFirestoreClient はfirestore clientがあればcloseします
func (f *Firebase) FinishFirestoreClient(cli *firestore.Client) error {
	if cli == nil {
		return nil
	}
	return cli.Close()
}

// NewAuthClint is return auth client
func (f *Firebase) NewAuthClint(ctx context.Context) (*auth.Client, error) {
	app, err := f.newApp(ctx)
	if err != nil {
		return nil, err
	}
	return app.Auth(ctx)
}
