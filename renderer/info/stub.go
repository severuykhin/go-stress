package info

import "github.com/severuykhin/go-stress/interfaces"

type StubInfoRenderer struct {
}

func NewStubResultRenderer() *StubInfoRenderer {
	return &StubInfoRenderer{}
}

func (srr *StubInfoRenderer) Render(optFunc ...interfaces.InfoOptFunc) {
	// do nothing
}
