package repository

import (
	"context"
)

// OtherService は同一PJ内の他GAEサービスです
// サービスに合わせて「OtherService」の名称を変更してください
type OtherService interface {
	Ping(ctx context.Context) ([]byte, error)
}