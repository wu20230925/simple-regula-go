package regula

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wu20230925/simple-regula-go/model"
)

const (
	basePath = "https://api.regulaforensics.com"
)

func TestClient_Ping(t *testing.T) {
	ping, err := NewClient(WithBaseURL(basePath)).Ping(context.Background(), model.PingParams{})
	assert.Equal(t, nil, err)
	t.Log(ping)
}

func TestClient_ApiProcess(t *testing.T) {
	processResult, err := NewClient(WithBaseURL(basePath)).ApiProcess(
		context.Background(),
		model.ApiProcessParams{},
		model.ProcessRequest{
			ProcessParam: model.ProcessParams{
				Scenario: model.ScenarioFULLPROCESS,
				ResultTypeOutput: &[]model.Result{
					model.ResultTEXT, model.ResultSTATUS,
				},
			},
		})
	assert.Equal(t, nil, err)

	parser := NewProcessResultParser(processResult)
	exist, textItem := parser.GetByResultType(model.ResultTEXT)
	if !exist {
		t.Fatal("not exist text")
	}

	result, err := textItem.AsTextDataResult()
	if err != nil {
		return
	}
	t.Log(result)
}
