package progress

import (
	"context"
	"fmt"
	"time"
)

type StdOutTimeBarProgressRenderer struct {
}

func NewStdOutTimeBarProgressRenderer() *StdOutTimeBarProgressRenderer {
	return &StdOutTimeBarProgressRenderer{}
}

func (sotb *StdOutTimeBarProgressRenderer) Run(ctx context.Context, target int) {

	progressDurationStep := time.Millisecond * 1
	progressTotalDuration := time.Duration(0)
	progressTarget := time.Millisecond * time.Duration(target*1000)

	for {
		stepStart := time.Now()
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			return
		default:
			time.Sleep(progressDurationStep)
			stepEnd := time.Since(stepStart)
			fmt.Printf("\033[2K\r%s", sotb.buildProgressString(int(progressTarget), int(progressTotalDuration)))
			progressTotalDuration += stepEnd
			continue
		}
	}
}

func (sotd *StdOutTimeBarProgressRenderer) buildProgressString(target, current int) string {
	if current > target {
		current = target
	}
	complete := (current / (target / 100)) / 2

	// start := fmt.Sprintf("%ds--", time.Duration(current)/time.Second)
	str := ""
	// str += start

	for i := 0; i < complete; i++ {
		str += "="
	}

	str += ">"

	for i := complete; i < 50; i++ {
		str += "-"
	}
	return str
}
