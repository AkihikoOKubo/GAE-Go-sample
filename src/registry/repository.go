package registry

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/repository"
	"github.com/AkihikoOkubo/gae-go-sample/src/infra/firebase"
	"github.com/AkihikoOkubo/gae-go-sample/src/infra/other"
)


// Repository はRepositoryの生成関数を定義します
// Usecaseの生成時にリポジトリの実装をDIするために使います
type Repository interface {
	NewFirebase() repository.Firebase
	NewOtherService() repository.OtherService
}

// RepositoryImpl is ...
type RepositoryImpl struct {
	Cnf *config.ServerConfig
}

// NewRepository はRepositoryを返します
func NewRepository(cnf *config.ServerConfig) Repository {
	return &RepositoryImpl{
		Cnf: cnf,
	}
}

// NewFirebase はFirebaseリポジトリを返します
func (r *RepositoryImpl) NewFirebase() repository.Firebase {
	return &firebase.Firebase{
		Cnf: &r.Cnf.Firebase,
	}
}

// NewOtherService は他サービスの呼び出しリポジトリを返します
func (r *RepositoryImpl) NewOtherService() repository.OtherService {
	return other.NewClient(&r.Cnf.OtherService)
}
