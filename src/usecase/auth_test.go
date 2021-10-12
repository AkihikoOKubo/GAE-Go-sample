package usecase

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/model"
	"github.com/AkihikoOkubo/gae-go-sample/src/domain/repository/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Auth", func() {

	var au *AuthUsecase

	var (
		ctx      = context.Background()
		cnf      = &config.ServerConfig{}
		mockCtrl *gomock.Controller
		fb       *mock_repository.MockFirebase
		token *model.Token
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		fb = mock_repository.NewMockFirebase(mockCtrl)

		au = &AuthUsecase{
			cnf:      cnf,
			fb:       fb,
		}

		token = &model.Token{
			Token:     "fb_token",
			ClientID: "ClientID",
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	authCli := &auth.Client{}

	Describe("VerifyToken", func() {
		Context("Admin", func() {
			Context("有効なTokenである場合", func() {
				It("nilが返る", func() {
					fb.EXPECT().NewAuthClint(ctx).Return(authCli, nil)
					fb.EXPECT().VerifyIDToken(ctx, authCli, token).Return(nil)
					err := au.VerifyToken(ctx, token)
					Ω(err).To(BeNil())
				})
			})
			Context("無効なTokenである場合", func() {
				It("エラーが返る", func() {
					fb.EXPECT().NewAuthClint(ctx).Return(authCli, nil)
					fb.EXPECT().VerifyIDToken(ctx, authCli, token).Return(config.ErrUnauthorized)
					err := au.VerifyToken(ctx, token)
					Ω(err).To(Equal(config.ErrUnauthorized))
				})
			})
		})
	})
})
