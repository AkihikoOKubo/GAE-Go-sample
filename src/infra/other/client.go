package other

import (
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/request"
)

// Client は通信クライアントです
type Client struct {
	request.Client
	Cnf *config.OtherService
}

// NewClient is return client
func NewClient(cnf *config.OtherService) *Client {
	return NewClientWithConfig(cnf)
}

// NewClientWithConfig is return client
func NewClientWithConfig(cnf *config.OtherService) *Client {
	cli := &Client{Cnf: cnf}
	cli.Client.Host = cnf.Host
	cli.Client.Audience = cnf.Audience
	return cli
}