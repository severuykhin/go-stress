package interfaces

import (
	"net/http"
	"time"

	"github.com/severuykhin/go-stress/structs"
)

type ConfigParserInterface interface {
	Parse() ([]structs.Stage, error)
}

/*
	Отображение информации
*/

type InfoRenderer interface {
	Render(structs.InfoData)
}

type DataCollection interface {
	CollectError(requestElapsedTime time.Duration, err error)
	CollectRequest(requestElapsedTime time.Duration, response *http.Response)
	GetCommonReport() structs.InfoData
	GetDetailedReport() structs.InfoData
}

type DataCollectionFactory interface {
	Create(config structs.CollectionConfig) DataCollection
}

// Render(WithHeader([]string), WithData([][]string))
