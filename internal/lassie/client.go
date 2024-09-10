package lassie

import (
	"context"
	"io"

	"github.com/filecoin-project/go-jsonrpc" // Ensure this import is correct
	"github.com/filecoin-project/lassie/api"
)

type Client struct {
	api api.LassieAPI
}

func NewClient(apiAddr string) (*Client, error) {
	var lassieAPI api.LassieAPI
	closer, err := jsonrpc.NewMergeClient(context.Background(), apiAddr, "Filecoin", []interface{}{&lassieAPI}, nil)
	if err != nil {
		return nil, err
	}
	_ = closer

	return &Client{api: lassieAPI}, nil
}

func (c *Client) Retrieve(cid string) (io.Reader, error) {
	return c.api.Retrieve(context.Background(), cid)
}
