package input

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/thedevsaddam/govalidator"
)

// User はPOST bodyをマッピングする構造体です
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserRules はPOST bodyのバリデーションルールです
var UserRules = govalidator.MapData{
	"id":   []string{"required"},
	"name": []string{"required"},
	"age":  []string{"required"},
}

// ToModel はモデルを作成します
func (u *User) ToModel() *model.User {
	usr := &model.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
	}

	return usr
}
