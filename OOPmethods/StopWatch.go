package OOPmethods

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	StartTime time.Time
	Splits    []time.Duration
}

func (s *Stopwatch) Start() {
	s.StartTime = time.Now()
}

func (s *Stopwatch) SaveSplit() {
	duration := time.Since(s.StartTime)
	s.Splits = append(s.Splits, duration)
}

func (s *Stopwatch) GetResults() {
	fmt.Println(s.Splits)
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	sw.GetResults()
}
