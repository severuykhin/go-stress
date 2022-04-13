package app

import (
	"context"
	"io"
	"net/http"
)

type Client interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

/*
	Отображение общего прогресса выполнения
	target - целевое значение которым может быть выражен общий прогресс выполнения
*/
type ProgressRenderer interface {
	Run(ctx context.Context, target int)
}

/*
	Отображение результатов тестирования
	Агрегированные результаты тестирования предоставляются в виде нескольких массиво строк
*/
type ResultRenderer interface {
	Render(results [][]string)
}
