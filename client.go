package vscale

import (
	"net/http"
	"os"
	"time"
)

type Client struct {
	Token   string
	client  *http.Client
	BaseURL string
	Region  RegionsService
	Scalet  ScaletsService
}

func NewClient(token string) *Client {
	client := &Client{
		Token: token,
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &loggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
		BaseURL: "https://api.vscale.io",
	}

	client.Region = &RegionsServiceOp{client: client}
	client.Scalet = &ScaletsServiceOp{client: client}

	return client
}
