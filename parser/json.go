package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/severuykhin/go-stress/structs"
)

type jsonParser struct {
	fileName string
}

func NewJsonParser(fileName string) *jsonParser {
	return &jsonParser{
		fileName: fileName,
	}
}

func (jp *jsonParser) Parse() ([]structs.Stage, error) {

	data, err := os.ReadFile(jp.fileName)

	if err != nil {
		return nil, err
	}

	var stages []structs.Stage

	err = json.Unmarshal(data, &stages)

	if err != nil {
		return nil, err
	}

	for i, stage := range stages {
		if err := stage.IsValid(); err != nil {
			return nil, fmt.Errorf("error: stage %d: %s", i, err)
		}
	}

	return stages, nil
}
