package usecase

import (
	"context"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/repository"
	"github.com/AkihikoOkubo/gae-go-sample/src/registry"
)

// AuthUsecaseImpl は認証・認可ユースケースの実装です
type AuthUsecaseImpl interface {
	VerifyToken(ctx context.Context, t *model.Token) error
}

// AuthUsecase は認証・認可ユースケースです
type AuthUsecase struct {
	cnf      *config.ServerConfig
	fb       repository.Firebase
}

// NewAuthUsecase はユースケースを作成します
func NewAuthUsecase(cnf *config.ServerConfig, repo registry.Repository) AuthUsecaseImpl {
	return &AuthUsecase{
		cnf:      cnf,
		fb:       repo.NewFirebase(),
	}
}

// VerifyToken はauth tokenを検証します
// tokenのタイプ(line || firebase)に合わせて認証処理を行います
func (au *AuthUsecase) VerifyToken(ctx context.Context, t *model.Token) error {
	cli, err := au.fb.NewAuthClint(ctx)
	if err != nil {
		return err
	}
	return au.fb.VerifyIDToken(ctx, cli, t)
}