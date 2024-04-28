package regula

import (
	"encoding/json"
	"log"

	"github.com/wu20230925/simple-regula-go/model"
)

type ProcessResultParser struct {
	processResult model.ProcessResponse

	models map[model.Result]model.ContainerList_List_Item
}

func NewProcessResultParser(response model.ProcessResponse) ProcessResultParser {
	parser := ProcessResultParser{
		processResult: response,
	}
	parser.process()
	return parser
}

func (p *ProcessResultParser) process() {
	if p.processResult.ProcessingFinished != model.FINISHED {
		return
	}
	p.models = make(map[model.Result]model.ContainerList_List_Item)
	var itemType struct {
		ResultType model.Result `json:"result_type"`
	}
	for _, item := range p.processResult.ContainerList.List {
		itemData, err := item.MarshalJSON()
		if err != nil {
			log.Printf("process ContainerList item err:%v", err)
			return
		}
		if err = json.Unmarshal(itemData, &itemType); err != nil {
			log.Printf("process ContainerList item err:%v", err)
			return
		}
		p.models[itemType.ResultType] = item
	}
}

func (p *ProcessResultParser) GetByResultType(resultType model.Result) (bool, model.ContainerList_List_Item) {
	val, ok := p.models[resultType]
	return ok, val
}
