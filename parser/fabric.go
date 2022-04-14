package parser

import (
	"fmt"
	"regexp"

	"github.com/severuykhin/go-stress/interfaces"
	"github.com/severuykhin/go-stress/pkg/arrays"
)

var FILE_FLAG_NAMES = []string{"-f", "--file"}

func CreateFrom(args []string) (interfaces.ConfigParserInterface, error) {
	filePath := ""

	fileNameRegex, err := regexp.Compile(`^[a-zA-Z]{0,}\.json$`)

	if err != nil {
		return nil, err
	}

	for _, flagName := range FILE_FLAG_NAMES {
		found := arrays.FindString(flagName, args) >= 0
		if found {
			fileNameIndex := arrays.FindStringByRegex(fileNameRegex, args)
			if fileNameIndex >= 0 {
				filePath = args[fileNameIndex]
			} else {
				return nil, fmt.Errorf("invalid file name. the file name must satisfy the condition ^[a-zA-Z]{0,}.json$")
			}
		}
	}

	if filePath != "" {
		return NewJsonParser(filePath), nil
	} else {
		return NewCliParser(args), nil
	}
}
