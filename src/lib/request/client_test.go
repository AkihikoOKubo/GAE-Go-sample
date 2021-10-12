package request

import (
	"context"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("Api", func() {
	var s *ghttp.Server
	var c *Client

	BeforeEach(func() {
		s = ghttp.NewServer()
		c = &Client{
			Client: new(http.Client),
			Host:   s.URL(),
			IsTest: true,
		}
	})

	AfterEach(func() {
		s.Close()
	})

	path := "/some/path"
	target := struct {
		Name string
	}{}

	Describe("GetJSON", func() {
		It("Requestを送ることができる", func() {
			s.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest(http.MethodGet, path),
					ghttp.RespondWith(http.StatusOK, []byte(`{"name": "hoge"}`)),
				))

			err := c.GetJSON(context.Background(), path, nil, &target)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(target).Should(Equal(struct {
				Name string
			}{
				Name: "hoge",
			}))
		})
	})

	Describe("PostJSON", func() {
		type data struct {
			ID string
		}

		Context("StatusNoContentの場合", func() {
			It("Requestを送ることができる", func() {
				s.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyHeader(http.Header{
							"Content-Type": []string{"application/json"},
						}),
						ghttp.VerifyRequest(http.MethodPost, path),
						ghttp.VerifyJSONRepresenting(data{ID: "test"}),
						ghttp.RespondWith(http.StatusNoContent, nil),
					))

				err := c.PostJSON(context.Background(), path, nil, []byte(`{"ID": "test"}`), nil)
				Ω(err).ShouldNot(HaveOccurred())
			})

		})

		Context("StatusOKの場合", func() {
			It("Requestを送ることができる", func() {
				s.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyHeader(http.Header{
							"Content-Type": []string{"application/json"},
						}),
						ghttp.VerifyRequest(http.MethodPost, path),
						ghttp.VerifyJSONRepresenting(data{ID: "test"}),
						ghttp.RespondWith(http.StatusOK, []byte(`{"name": "hoge"}`)),
					))

				err := c.PostJSON(context.Background(), path, nil, []byte(`{"ID": "test"}`), &target)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(target).Should(Equal(struct {
					Name string
				}{
					Name: "hoge",
				}))
			})
		})
	})
})
