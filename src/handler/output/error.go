package output

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
)

// ErrResponse is error response object
type ErrResponse struct {
	Message string `json:"message"`
}

// NewErrResponse はWrapErrを元にエラーレスポンスを作成します
func NewErrResponse(err model.WrapErr) *ErrResponse {
	return &ErrResponse{
		Message: err.Message,
	}
}
