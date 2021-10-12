package bridge

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/registry"
	"github.com/AkihikoOkubo/gae-go-sample/src/usecase"
)

// Handler is ...
type Handler struct {
	Cnf *config.ServerConfig
	uc  usecase.BridgeUsecaseImpl
}

// NewHandler はHandlerを返します
func NewHandler(cnf *config.ServerConfig) *Handler {
	return &Handler{
		Cnf: cnf,
		uc:  usecase.NewBridgeUsecase(cnf, registry.NewRepository(cnf)),
	}
}


