package app

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Clients  int
	Url      string
	Duration int // sec - общая продолжительность тестирования
	Timeout  int
	Method   string
}

type ResultRow struct {
	Type  string
	Count int
	Min   time.Duration
	Max   time.Duration
	Avg   time.Duration
}

type app struct {
	config           Config
	client           Client
	progressRenderer ProgressRenderer
	resultRenderer   ResultRenderer
}

func New(config Config, client Client, progressRenderer ProgressRenderer, resultRenderer ResultRenderer) *app {
	return &app{
		config:           config,
		client:           client,
		progressRenderer: progressRenderer,
		resultRenderer:   resultRenderer,
	}
}

func (a *app) Run() {
	var wg sync.WaitGroup
	var totalRequets int
	var totalElapsedTime time.Duration
	var mux sync.Mutex

	totalResultRow := ResultRow{
		Min: time.Duration(0),
		Max: time.Duration(0),
		Avg: time.Duration(0),
	}

	errorResultRow := ResultRow{
		Count: 0,
		Type:  "error",
	}

	requestAttemptTimeout := time.Millisecond * time.Duration(a.config.Timeout)
	totalStressDuration := time.Second * time.Duration(a.config.Duration)

	for i := 0; i < a.config.Clients; i++ {
		wg.Add(1)

		go func(tsd time.Duration, responseAttemptTimeout time.Duration) {

			minRequestElapsedTime := time.Duration(0)
			maxRequestElapsedTime := time.Duration(0)

			// общее время затраченное клиентом
			clientStressDuration := time.Duration(0)
			clientTotalRequests := 0
			clientTotalRequestsTime := time.Duration(0)
			clientTotalRequestDuration := time.Duration(0)

			for {
				// Время начала запроса
				requestStartTime := time.Now()
				_, err := a.client.Get(a.config.Url)
				if err != nil {
					fmt.Println(err)
					mux.Lock()
					errorResultRow.Count += 1
					mux.Unlock()
				} else {
					// todo
				}

				// Временная метка окончания запроса
				requestEndTime := time.Now()
				// Время затраченное на запрос
				requestElapsedTime := requestEndTime.Sub(requestStartTime)

				if requestElapsedTime > maxRequestElapsedTime {
					maxRequestElapsedTime = requestElapsedTime
				}

				if minRequestElapsedTime == 0 || (requestElapsedTime < minRequestElapsedTime) {
					minRequestElapsedTime = requestElapsedTime
				}

				totalResponseElapsedTime := requestElapsedTime + responseAttemptTimeout
				clientTotalRequestDuration += totalResponseElapsedTime
				clientTotalRequestsTime += requestElapsedTime

				time.Sleep(responseAttemptTimeout)

				clientStressDuration += totalResponseElapsedTime
				clientTotalRequests++

				if clientStressDuration >= tsd {
					break
				}
			}

			mux.Lock()
			totalRequets += clientTotalRequests
			totalElapsedTime += clientTotalRequestsTime
			mux.Unlock()

			mux.Lock()
			if maxRequestElapsedTime > totalResultRow.Max {
				totalResultRow.Max = maxRequestElapsedTime
			}
			if totalResultRow.Min == 0 {
				totalResultRow.Min = minRequestElapsedTime
			} else if minRequestElapsedTime < totalResultRow.Min {
				totalResultRow.Min = minRequestElapsedTime
			}
			mux.Unlock()

			wg.Done()

		}(totalStressDuration, requestAttemptTimeout)
	}

	ctx, cancelProgress := context.WithCancel(context.TODO())

	go a.progressRenderer.Run(ctx, a.config.Duration)

	wg.Wait()
	cancelProgress()

	results := [][]string{
		{
			"all",
			strconv.Itoa(totalRequets),
			strconv.Itoa(totalRequets/a.config.Duration) + " req/s",
			strconv.Itoa(int(totalResultRow.Min)/int(time.Millisecond)) + "ms",
			strconv.Itoa(int(totalResultRow.Max)/int(time.Millisecond)) + "ms",
			strconv.Itoa((int(totalElapsedTime)/int(time.Millisecond))/totalRequets) + "ms",
		},
		{
			"ok",
			strconv.Itoa(totalRequets - errorResultRow.Count),
		},
		{
			"errors",
			strconv.Itoa(errorResultRow.Count),
		},
	}

	a.resultRenderer.Render(results)
}
