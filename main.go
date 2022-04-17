package main

import (
	"log"
	"os"

	"github.com/severuykhin/go-stress/app"
	"github.com/severuykhin/go-stress/client"
	"github.com/severuykhin/go-stress/collection"
	"github.com/severuykhin/go-stress/parser"
	"github.com/severuykhin/go-stress/renderer/info"
	"github.com/severuykhin/go-stress/renderer/progress"
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

	parser, err := parser.CreateFrom(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	stages, err := parser.Parse()

	if err != nil {
		log.Fatal(err)
	}

	progressRenderer := progress.NewStdOutTimeBarProgressRenderer()
	infoRederer := info.NewTableViewInfoRenderer()
	appClient := client.NewHttpClient()
	dataCollectionFactory := collection.New()
	application := app.New(appClient, progressRenderer, infoRederer, dataCollectionFactory)

	application.Run(stages)

}
