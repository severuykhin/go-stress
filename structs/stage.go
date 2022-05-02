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

func (s *Stage) GetValuesFormatted() []string {
	return []string{
		s.Url,
		s.Method,
		strconv.Itoa(s.Duration) + "s",
		strconv.Itoa(s.Timeout) + "ms",
		strconv.Itoa(s.Clients) + "vc",
	}
}

func (s *Stage) IsValid() error {
	if s.Url == "" {
		return fmt.Errorf("url must be not empty")
	}

	return nil
}
