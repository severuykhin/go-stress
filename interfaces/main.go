package interfaces

import "github.com/severuykhin/go-stress/structs"

type ConfigParserInterface interface {
	Parse() ([]structs.Stage, error)
}

/*
	Отображение информации
*/

type InfoRenderer interface {
	Render(structs.InfoData)
}

// Render(WithHeader([]string), WithData([][]string))
