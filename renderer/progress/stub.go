package progress

import "context"

type StubProgressRenderer struct {
}

func NewStubProgressRenderer() *StubProgressRenderer {
	return &StubProgressRenderer{}
}

func (spr *StubProgressRenderer) Run(ctx context.Context, target int) {
	// do nothing
}
