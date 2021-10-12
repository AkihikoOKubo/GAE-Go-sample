package config

import (
	"github.com/kelseyhightower/envconfig"
)

// ServerConfig は環境変数を提供する構造体です。
type ServerConfig struct {
	App App
	Firebase     Firebase
	OtherService OtherService
}

// App はアプリケーション全体の設定です
type App struct {
	ProjectID string // GCPのProjectIDです
}

// Firebase はFirebaseと接続するための設定です
type Firebase struct {
	ProjectID string // GCPのProjectIDです
}

// OtherService は同一PJ内に他GAEサービスがあると仮定した接続設定です
// IAP環境下でHTTPリクエストでサービス間通信する場合のサンプルコードとなっています
type OtherService struct {
	ProjectID string // GCPのProjectIDです
	Host      string // 接続先のHOST(eg. https://${sercice_name}-${project_name}.appspot.com)
	Audience  string // IAPを有効化すると発行されるCLIENT_IDを入れます
}

// yaml上の定数のPrefix(キャメルケースに変換されます)
const prefix = "app"

var cnf *ServerConfig

// NewConfig はServerConfigを返します
// app_{ENV}.yamlから環境変数を取得してセットします
func NewConfig() (*ServerConfig, error) {
	if cnf != nil {
		return cnf, nil
	}

	cnf = &ServerConfig{}
	if err := envconfig.Process(prefix, cnf); err != nil {
		return nil, err
	}

	return cnf, nil
}
