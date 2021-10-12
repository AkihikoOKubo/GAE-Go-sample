package usecase

import (
	"context"
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/repository"
	"github.com/AkihikoOkubo/gae-go-sample/src/registry"
	"time"
)

// UserUsecaseImpl はUser関連の実装を持ちます.
type UserUsecaseImpl interface {
	FindByID(ctx context.Context, id string) (*model.User, error)
	Register(ctx context.Context, user *model.User, now time.Time) error
}

// UserUsecase はUser関連のユースケースです.
type UserUsecase struct {
	cnf *config.ServerConfig
	fb  repository.Firebase
}

// NewUserUsecase はユースケースを作成します
func NewUserUsecase(cnf *config.ServerConfig, repo registry.Repository) UserUsecaseImpl {
	return &UserUsecase{
		cnf: cnf,
		fb:  repo.NewFirebase(),
	}
}

// FindByID はUserIDを指定してUserを取得します
func (u *UserUsecase) FindByID(ctx context.Context, id string) (*model.User, error) {
	cli, err := u.fb.NewFirestoreClint(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if er := u.fb.FinishFirestoreClient(cli); er != nil {
			fmt.Print(err.Error())
		}
	}()
	return u.fb.User(ctx, cli, id)
}

// Register はUserを登録します
func (u *UserUsecase) Register(ctx context.Context, user *model.User, now time.Time) error {
	cli, err := u.fb.NewFirestoreClint(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if er := u.fb.FinishFirestoreClient(cli); er != nil {
			fmt.Print(err.Error())
		}
	}()
	user.UpdatedAt = now
	return u.fb.SaveUser(ctx, cli, user)
}