package spin

import "fmt"

type Spinner struct {
}

type Stats struct {
}

func (s *Spinner) Spin() {
	s.start()
}

func (s *Spinner) start() {
	fmt.Println("spinner started")
}

func New() *Spinner {
	return &Spinner{}
}
