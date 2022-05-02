package app

import (
	"context"
	"sync"
	"time"

	"github.com/severuykhin/go-stress/interfaces"
	"github.com/severuykhin/go-stress/structs"
)

type app struct {
	client                Client
	progressRenderer      ProgressRenderer
	infoRenderer          interfaces.InfoRenderer
	dataCollectionFactory interfaces.DataCollectionFactory
}

func New(
	client Client,
	progressRenderer ProgressRenderer,
	infoRenderer interfaces.InfoRenderer,
	dataCollectionFactory interfaces.DataCollectionFactory,
) *app {
	return &app{
		client:                client,
		progressRenderer:      progressRenderer,
		infoRenderer:          infoRenderer,
		dataCollectionFactory: dataCollectionFactory,
	}
}

func (a *app) Run(stages []structs.Stage) {
	for _, stage := range stages {
		a.RunStage(stage)
	}
}

func (a *app) RunStage(stage structs.Stage) {
	a.infoRenderer.Render(
		structs.InfoData{
			Header: stage.GetFields(),
			Data: [][]string{
				stage.GetValuesFormatted(),
			},
		},
	)

	stageDataCollection := a.dataCollectionFactory.Create(structs.CollectionConfig{
		RequstTimeout: time.Millisecond * time.Duration(stage.Timeout),
		Duration:      time.Second * time.Duration(stage.Duration),
	})

	var wg sync.WaitGroup
	for i := 0; i < stage.Clients; i++ {
		wg.Add(1)

		go func(st structs.Stage) {

			requestAttemptTimeout := time.Millisecond * time.Duration(stage.Timeout)
			totalDuration := time.Second * time.Duration(stage.Duration)
			clientTotalDuration := time.Duration(0)

			for {
				// Время начала запроса
				requestStartTime := time.Now()
				response, err := a.client.Get(stage.Url)

				// Временная метка окончания запроса
				requestEndTime := time.Now()
				// Время затраченное на запрос
				requestElapsedTime := requestEndTime.Sub(requestStartTime)
				totalRequestElapsedTime := requestElapsedTime + requestAttemptTimeout

				if err != nil {
					stageDataCollection.CollectError(requestElapsedTime, err)
				} else {
					stageDataCollection.CollectRequest(requestElapsedTime, response)
				}

				clientTotalDuration += totalRequestElapsedTime

				time.Sleep(requestAttemptTimeout)

				// если клиент потратил все время, отведенное на выполнение - остановим клинта
				if clientTotalDuration >= totalDuration {
					break
				}
			}

			wg.Done()
		}(stage)
	}

	ctx, cancel := context.WithCancel(context.TODO())
	go a.progressRenderer.Run(ctx, stage.Duration)

	wg.Wait()
	cancel()

	commonReport := stageDataCollection.GetCommonReport()
	a.infoRenderer.Render(commonReport)

	detailedReport := stageDataCollection.GetDetailedReport()
	a.infoRenderer.Render(detailedReport)
}
