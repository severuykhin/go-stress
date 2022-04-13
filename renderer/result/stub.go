package result

type StubResultRenderer struct {
}

func NewStubResultRenderer() *StubResultRenderer {
	return &StubResultRenderer{}
}

func (srr *StubResultRenderer) Render(results [][]string) {
	// do nothing
}
