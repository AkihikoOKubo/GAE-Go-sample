package output

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"time"
)

// User はユーザのレスポンス用構造体です
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"line_id"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser is return User
func NewUser(u *model.User) *User {
	if u == nil {
		return &User{}
	}

	return &User{
		ID:        u.ID,
		Name:      u.Name,
		Age:       u.Age,
		UpdatedAt: u.UpdatedAt,
	}
}
