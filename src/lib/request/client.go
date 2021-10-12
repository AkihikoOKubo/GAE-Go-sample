package request

import (
	"context"
	"fmt"
	"github.com/AkihikoOkubo/gae-go-sample/src/config"
	"github.com/AkihikoOkubo/gae-go-sample/src/lib/log"
	"github.com/pkg/errors"
	"google.golang.org/api/idtoken"
	"io/ioutil"
	"net/http"
)

// Client はHTTPクライアントです
type Client struct {
	Client   *http.Client
	Host     string
	Audience string
}

// Get is execute Get HTTP Request
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	var err error
	c.Client, err = idtoken.NewClient(ctx, c.Audience)
	if err != nil {
		log.Logger(ctx).Error(err.Error())
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.Host, path), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		log.Logger(ctx).Error(err.Error())
		return nil, config.ErrHTTPClientError
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Logger(ctx).Error(err.Error())
		}
	}()

	var bs []byte
	if res.StatusCode != http.StatusNoContent {
		bs, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
	}

	return bs, c.getError(res)
}

func (c *Client) getError(res *http.Response) error {
	if res.StatusCode < 400 {
		return nil
	}

	switch res.StatusCode {
	case http.StatusBadRequest:
		return config.ErrBadRequest
	case http.StatusUnauthorized:
		return config.ErrUnauthorized
	case http.StatusForbidden:
		return config.ErrForbidden
	case http.StatusNotFound:
		return config.ErrNotFound
	case http.StatusMethodNotAllowed:
		return config.ErrMethodNotAllowed
	case http.StatusInternalServerError:
		return config.ErrInternalServerError
	case http.StatusBadGateway:
		return config.ErrBadGateway
	case http.StatusServiceUnavailable:
		return config.ErrServiceUnavailable
	case http.StatusGatewayTimeout:
		return config.ErrGatewayTimeout
	default:
		return errors.New(http.StatusText(res.StatusCode))
	}
}
