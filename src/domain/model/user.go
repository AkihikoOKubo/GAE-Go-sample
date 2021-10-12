package model

import (
	"time"
)

// User はユーザの構造体です
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	UpdatedAt time.Time `json:"updated_at"`
}