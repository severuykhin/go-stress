package collection

import (
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/severuykhin/go-stress/structs"
)

type dataCollection struct {
	requestTimeout time.Duration
	duration       time.Duration

	totalRequestCount  int // сколько всего запросов совершено
	failedRequestCount int // сколько запросов не выполнено

	totalElapsedTime   time.Duration // общее затраченоне время на выполнение запросов
	avgRequestDuration time.Duration // среднее время потраченное на запрос
	minRequestDuration time.Duration // минимальное время затраченное на запрос
	maxRequestDuration time.Duration // максимальное время затраченное на запрос

	latencies []int

	mux sync.Mutex
}

func (dc *dataCollection) CollectError(requestElapsedTime time.Duration, err error) {
	dc.mux.Lock()
	dc.totalRequestCount += 1
	dc.failedRequestCount += 1
	dc.mux.Unlock()
}

func (dc *dataCollection) CollectRequest(requestElapsedTime time.Duration, response *http.Response) {
	dc.mux.Lock()
	dc.totalRequestCount += 1
	dc.totalElapsedTime += requestElapsedTime
	dc.avgRequestDuration = dc.totalElapsedTime / time.Duration(dc.totalRequestCount)
	dc.mux.Unlock()

	dc.mux.Lock()
	if dc.maxRequestDuration < requestElapsedTime {
		dc.maxRequestDuration = requestElapsedTime
	}
	dc.mux.Unlock()

	dc.mux.Lock()
	if dc.minRequestDuration == time.Duration(0) {
		dc.minRequestDuration = requestElapsedTime
	} else if requestElapsedTime < dc.minRequestDuration {
		dc.minRequestDuration = requestElapsedTime
	}
	dc.mux.Unlock()

	dc.mux.Lock()
	dc.latencies = append(dc.latencies, int(requestElapsedTime))
	dc.mux.Unlock()
}

/*

      | count | qps | min  | max   | avg  |
total | 100   | 10  | 10ms | 100ms | 50ms |
ok    | 90    |
error | 10    |

*/
func (dc *dataCollection) GetCommonReport() structs.InfoData {
	qps := int(float64(dc.totalRequestCount-dc.failedRequestCount) / dc.duration.Seconds())
	sort.Ints(dc.latencies)
	p75 := dc.getPMetrics(75).String()
	p99 := dc.getPMetrics(99).String()

	return structs.InfoData{
		Header: []string{
			"", "count", "min", "max", "avg", "qps", "p75", "p99",
		},
		Data: [][]string{
			{
				"ok",
				strconv.Itoa(dc.totalRequestCount - dc.failedRequestCount),
				"-",
				"-",
				"-",
				"-",
				"-",
			},
			{
				"fail",
				strconv.Itoa(dc.failedRequestCount),
				"-",
				"-",
				"-",
				"-",
				"-",
			},
			{
				"total",
				strconv.Itoa(dc.totalRequestCount),
				dc.minRequestDuration.String(),
				dc.maxRequestDuration.String(),
				dc.avgRequestDuration.String(),
				strconv.Itoa(qps) + "/s",
				p75,
				p99,
			},
		},
	}
}

func (dc *dataCollection) GetDetailedReport() structs.InfoData {
	return structs.InfoData{}
}

func (dc *dataCollection) getPMetrics(point int) time.Duration {

	res := 0
	pointReal := float32(point)
	percent := float32(dc.totalRequestCount) / float32(100)
	for index, value := range dc.latencies {
		currentPoint := float32(index) / percent
		if currentPoint > pointReal {
			res = value
			break
		}
	}

	return time.Duration(res)
}
