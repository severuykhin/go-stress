package collection

import (
	"github.com/severuykhin/go-stress/interfaces"
	"github.com/severuykhin/go-stress/structs"
)

type dataCollectionFactory struct {
}

func New() *dataCollectionFactory {
	return &dataCollectionFactory{}
}

func (dcf *dataCollectionFactory) Create(config structs.CollectionConfig) interfaces.DataCollection {
	return &dataCollection{
		requestTimeout: config.RequstTimeout,
		duration:       config.Duration,
	}
}
