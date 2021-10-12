package firebase

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

const pathUsers = "sample-users"

// User はUserを取得します
func (f *Firebase) User(ctx context.Context, cli *firestore.Client, id string) (*model.User, error) {
	data := &model.User{}
	var err error

	var dSnap *firestore.DocumentSnapshot
	dSnap, err = cli.Collection(pathUsers).Doc(id).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return nil, config.ErrNotFound
		}
		return nil, err
	}
	err = dSnap.DataTo(data)
	if err != nil {
		return nil, err
	}

	return data, err
}

// SaveUser はUserをUpsertします
func (f *Firebase) SaveUser(ctx context.Context, cli *firestore.Client, user *model.User) error {
	_, err := cli.Collection(pathUsers).Doc(strconv.Itoa(user.ID)).Set(ctx, user)
	return err
}