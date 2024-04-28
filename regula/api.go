package regula

import (
	"context"

	"github.com/wu20230925/simple-regula-go/model"
)

type API interface {
	// Ping
	// @Description: /api/ping
	Ping(ctx context.Context, params model.PingParams) (model.DeviceInfo, error)

	// ApiProcess
	// @Description: /api/process
	ApiProcess(ctx context.Context, params model.ApiProcessParams, body model.ProcessRequest) (model.ProcessResponse, error)

	// TODO
	// /api/v2/tag/{tag_id}
	// /api/v2/transaction/{transaction_id}/file
	// /api/v2/transaction/{transaction_id}/results
	// /api/v2/transaction/{transaction_id}
	// /api/v2/transaction/{transaction_id}/process
}
