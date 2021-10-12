package usecase

import (
	"context"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/repository"
	"github.com/AkihikoOkubo/gae-go-sample/src/registry"
)

// BridgeUsecaseImpl is
type BridgeUsecaseImpl interface {
	Ping(ctx context.Context) ([]byte, error)
}

// BridgeUsecase is
type BridgeUsecase struct {
	cnf *config.ServerConfig
	o   repository.OtherService
}

// NewBridgeUsecase はユースケースを作成します
func NewBridgeUsecase(cnf *config.ServerConfig, repo registry.Repository) BridgeUsecaseImpl {
	return &BridgeUsecase{
		cnf: cnf,
		o:   repo.NewOtherService(),
	}
}

func (u *BridgeUsecase) Ping(ctx context.Context) ([]byte, error) {
	return u.o.Ping(ctx)
}
