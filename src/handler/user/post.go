package user

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/input"
	"github.com/AkihikoOkubo/gae-go-sample/src/handler/output"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/log"
	"github.com/alioygur/gores"
	"github.com/thedevsaddam/govalidator"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

// Post はユーザーデータを 更新・作成 します
func (h *Handler) Post(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var user input.User
	opts := govalidator.Options{
		Request: r,
		Data:    &user,
		Rules:   input.UserRules,
	}

	if result := govalidator.New(opts).ValidateJSON().Encode(); result != "" {
		e := model.NewWrapErr(config.ErrBadRequest)
		gores.JSON(w, e.StatusCode, output.NewErrResponse(e))
		return
	}

	err := h.Usecase.Register(ctx, user.ToModel(),time.Now())
	if err != nil {
		log.Logger(ctx).Error(err.Error())
		e := model.NewWrapErr(err)
		gores.JSON(w, e.StatusCode, output.NewErrResponse(e))
		return
	}

	gores.JSON(w, http.StatusNoContent, nil)
}