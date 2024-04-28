package regula

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/wu20230925/simple-regula-go/model"
)

type JsonUnmarshalFunc func(rsp *resty.Response, v interface{}) error

func JsonUnmarshalHandler(rsp *resty.Response, v interface{}) error {
	if rsp.StatusCode() != http.StatusOK {
		return fmt.Errorf("response status code: %s", rsp.Status())
	}
	return json.Unmarshal(rsp.Body(), &v)
}

type Client struct {
	cli     *resty.Client
	license string
}

func NewClient(opts ...Option) *Client {
	client := &Client{
		cli: resty.New(),
	}
	for _, opt := range opts {
		opt(client)
	}
	return client
}

func (c Client) Ping(ctx context.Context, params model.PingParams) (model.DeviceInfo, error) {
	var data model.DeviceInfo

	req := c.cli.R()
	if params.XRequestID != nil {
		req.SetHeader("X-RequestID", *params.XRequestID)
	}
	resp, err := req.Get("/api/ping")
	if err != nil {
		return data, err
	}
	if err = JsonUnmarshalHandler(resp, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (c Client) ApiProcess(ctx context.Context, params model.ApiProcessParams, body model.ProcessRequest) (model.ProcessResponse, error) {
	var data model.ProcessResponse

	req := c.cli.R().SetBody(body)
	if params.XRequestID != nil {
		req.SetHeader("X-RequestID", *params.XRequestID)
	}
	resp, err := req.Post("/api/process")
	if err != nil {
		return data, err
	}
	if err = JsonUnmarshalHandler(resp, &data); err != nil {
		return data, err
	}
	return data, nil
}
