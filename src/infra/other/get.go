package other

import (
	"context"
)

func (c *Client) Ping(ctx context.Context) ([]byte, error) {
	return c.Get(ctx, "/ping")
}