package main

import (
	"testing"

	"github.com/severuykhin/go-stress/app"
	"github.com/severuykhin/go-stress/client"
	"github.com/severuykhin/go-stress/renderer/progress"
	"github.com/severuykhin/go-stress/renderer/result"
)

func BenchmarkSample(b *testing.B) {
	config := app.Config{
		Duration: 5,
		Timeout:  70,
		Clients:  3,
		Url:      "http://api.app.loc",
		Method:   "GET",
	}

	progressRenderer := progress.NewStubProgressRenderer()
	resultRederer := result.NewStubResultRenderer()
	appClient := client.NewHttpClient()
	application := app.New(config, appClient, progressRenderer, resultRederer)

	for i := 0; i < b.N; i++ {
		application.Run()
	}
}
