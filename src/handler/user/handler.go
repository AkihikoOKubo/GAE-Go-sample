package user

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/registry"
	"github.com/AkihikoOkubo/gae-go-sample/src/usecase"
)

// Handler はUser関連の処理を持ちます
type Handler struct {
	Cnf     *config.ServerConfig
	Usecase usecase.UserUsecaseImpl
}

// NewHandler はHandlerを返します
func NewHandler(cnf *config.ServerConfig) *Handler {
	return &Handler{
		Cnf:     cnf,
		Usecase: usecase.NewUserUsecase(cnf, registry.NewRepository(cnf)),
	}
}
