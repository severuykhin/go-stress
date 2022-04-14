package structs

import (
	"fmt"
	"strconv"
)

type Stage struct {
	Clients  int
	Url      string
	Duration int // sec - общая продолжительность тестирования
	Timeout  int
	Method   string
}

func (s *Stage) GetFields() []string {
	return []string{
		"Url",
		"Method",
		"Duration",
		"Timeout",
		"Clients",
	}
}

func (s *Stage) GetValues() []string {
	return []string{
		s.Url,
		s.Method,
		strconv.Itoa(s.Duration),
		strconv.Itoa(s.Timeout),
		strconv.Itoa(s.Clients),
	}
}

func (s *Stage) IsValid() error {
	if s.Url == "" {
		return fmt.Errorf("url must be not empty")
	}

	return nil
}
