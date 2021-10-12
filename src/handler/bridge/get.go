package bridge

import (
	"context"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/output"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/log"
	"github.com/alioygur/gores"
	"net/http"
)

// PingOtherService は他サービスにpingを投げます
func (h *Handler) PingOtherService(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	bs, err := h.uc.Ping(ctx)
	if err != nil {
		log.Logger(ctx).Error(err.Error())
		e := model.NewWrapErr(err)
		gores.JSON(w, e.StatusCode, output.NewErrResponse(e))
		return
	}

	gores.JSON(w, http.StatusOK, string(bs))
}