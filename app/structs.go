package app

type Stage struct {
	Clients  int
	Url      string
	Duration int // sec - общая продолжительность тестирования
	Timeout  int
	Method   string
}
