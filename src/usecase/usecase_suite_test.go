package usecase

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestUsecase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Usecase Suite")
}
