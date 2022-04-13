package main

import (
	"flag"
	"log"

	"github.com/severuykhin/go-stress/app"
	"github.com/severuykhin/go-stress/client"
	"github.com/severuykhin/go-stress/renderer/progress"
	"github.com/severuykhin/go-stress/renderer/result"
)

/*

Command line arguments

-m --method  - GET,POST,DELETE,PUT,PATCH
-d --duration - продолжительность в секундах
-t --timeout - таймаут между запросами
-u --url     - http://some.site
-c --clients - 1 | 10 | 100

*/

func main() {

	var config app.Config

	if err := ParseCommandLineArguments(&config); err != nil {
		log.Fatal(err)
	}

	progressRenderer := progress.NewStdOutTimeBarProgressRenderer()
	resultRederer := result.NewTableViewResultRenderer()
	appClient := client.NewHttpClient()
	application := app.New(config, appClient, progressRenderer, resultRederer)

	application.Run()

}

func ParseCommandLineArguments(config *app.Config) error {
	clients := flag.Int("c", 1, "# of virtual clients")
	method := flag.String("m", "GET", "request method")
	url := flag.String("u", "", "request method")
	duration := flag.Int("d", 1, "stress test duration in seconds")
	timeout := flag.Int("t", 1000, "request timeout in milliseconds")
	flag.Parse()

	config.Clients = *clients
	config.Duration = *duration
	config.Url = *url
	config.Method = *method
	config.Timeout = *timeout

	return nil
}
