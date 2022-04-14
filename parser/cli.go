package parser

import (
	"flag"

	"github.com/severuykhin/go-stress/structs"
)

type CliParser struct {
	args []string
}

func NewCliParser(args []string) *CliParser {
	return &CliParser{
		args: args,
	}
}

func (cp *CliParser) Parse() ([]structs.Stage, error) {
	flagSet := flag.NewFlagSet("gostress-cli-args", flag.ContinueOnError)

	clients := flagSet.Int("c", 1, "# of virtual clients")
	method := flagSet.String("m", "GET", "request method")
	url := flagSet.String("u", "", "request url")
	duration := flagSet.Int("d", 1, "stress test duration in seconds")
	timeout := flagSet.Int("t", 1000, "request timeout in milliseconds")

	err := flagSet.Parse(cp.args)

	if err != nil {
		return nil, err
	}

	stage := structs.Stage{
		Url:      *url,
		Method:   *method,
		Clients:  *clients,
		Duration: *duration,
		Timeout:  *timeout,
	}

	if err := stage.IsValid(); err != nil {
		return nil, err
	}

	return []structs.Stage{
		stage,
	}, nil
}
