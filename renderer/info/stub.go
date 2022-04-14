package info

import (
	"github.com/severuykhin/go-stress/structs"
)

type StubInfoRenderer struct {
}

func NewStubResultRenderer() *StubInfoRenderer {
	return &StubInfoRenderer{}
}

func (srr *StubInfoRenderer) Render(infoData structs.InfoData) {
	// do nothing
}
