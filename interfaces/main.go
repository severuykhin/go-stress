package interfaces

import "github.com/severuykhin/go-stress/structs"

type ConfigParserInterface interface {
	Parse() ([]structs.Stage, error)
}

/*
	Отображение информации
*/

type InfoOptFunc func(options *structs.InfoOptions)

func WithData(data [][]string) InfoOptFunc {
	return func(options *structs.InfoOptions) {
		options.Data = data
	}
}

func WithHeader(header []string) InfoOptFunc {
	return func(options *structs.InfoOptions) {
		options.Header = header
	}
}

type InfoRenderer interface {
	Render(optFunc ...InfoOptFunc)
}

// Render(WithHeader([]string), WithData([][]string))
