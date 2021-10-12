package user

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/output"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/log"
	"github.com/alioygur/gores"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"
)

// Get はIDを指定してユーザを取得します
func (h *Handler) Get(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if id == "" {
		e := model.NewWrapErr(config.ErrBadRequest)
		gores.JSON(w, http.StatusBadRequest, output.NewErrResponse(e))
		return
	}

	u, err := h.Usecase.FindByID(ctx, id)
	if err != nil {
		log.Logger(ctx).Error(err.Error())
		e := model.NewWrapErr(err)
		gores.JSON(w, e.StatusCode, output.NewErrResponse(e))
		return
	}

	gores.JSON(w, http.StatusOK, output.NewUser(u))
}