package main

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"os"

	"github.com/wu20230925/simple-regula-go/model"
	"github.com/wu20230925/simple-regula-go/regula"
	"github.com/wu20230925/simple-regula-go/wrappers"
)

func main() {
	whitePage0, err := readFile("WHITE.jpg")
	if err != nil {
		log.Print(err)
		return
	}
	imageList := make([]model.ProcessRequestImage, 3)
	imageList[0] = model.ProcessRequestImage{
		ImageData: model.ImageData{Image: base64.StdEncoding.EncodeToString(whitePage0)},
		Light:     wrappers.Light(model.WHITE),
		PageIdx:   wrappers.Int(0),
	}
	irPage0, err := readFile("IR.jpg")
	if err != nil {
		log.Print(err)
		return
	}
	imageList[1] = model.ProcessRequestImage{
		ImageData: model.ImageData{Image: base64.StdEncoding.EncodeToString(irPage0)},
		Light:     wrappers.Light(model.IR),
		PageIdx:   wrappers.Int(0),
	}
	uvPage0, err := readFile("UV.jpg")
	if err != nil {
		log.Print(err)
		return
	}
	imageList[2] = model.ProcessRequestImage{
		ImageData: model.ImageData{Image: base64.StdEncoding.EncodeToString(uvPage0)},
		Light:     wrappers.Light(model.UV),
		PageIdx:   wrappers.Int(0),
	}

	processResult, err := regula.NewClient(regula.WithBaseURL("https://api.regulaforensics.com")).ApiProcess(
		context.Background(),
		model.ApiProcessParams{},
		model.ProcessRequest{
			ProcessParam: model.ProcessParams{
				Scenario: model.ScenarioFULLPROCESS,
				ResultTypeOutput: &[]model.Result{
					model.ResultTEXT, model.ResultSTATUS,
				},
			},
			List: &imageList,
		})
	if err != nil {
		log.Fatal(err)
		return
	}
	parser := regula.NewProcessResultParser(processResult)
	exist, textItem := parser.GetByResultType(model.ResultTEXT)
	if !exist {
		return
	}

	result, err := textItem.AsTextResult()
	if err != nil {
		return
	}
	log.Print(result)
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}
