package parser

import (
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
	return []structs.Stage{}, nil
}
