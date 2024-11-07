package queue

import "fmt"

func SimpleIterator(yield func(any) bool) {
	for i := 0; i < 100; i++ {
		if !yield(i) {
			return
		}
	}
}

type Queue struct {
	inputValues func(yield func(any) bool)
	input       chan any
	output      chan any
	steps       []queueStep
}

type queueStep struct {
	Input       chan any
	Output      chan any
	ProcessFunc func(input any) (any, error)
}

func NewQueue(inputValues func(yield func(any) bool)) Queue {
	var q Queue
	q.inputValues = inputValues
	q.input = make(chan any)
	q.output = q.input
	return q
}

func (q *Queue) AddStep(processFunc func(input any) (any, error)) {
	var s queueStep
	s.Output = make(chan any)
	s.ProcessFunc = processFunc
	l := len(q.steps)
	if l == 0 {
		s.Input = q.input
	} else {
		s.Input = q.steps[l-1].Output
	}
	q.output = s.Output
	q.steps = append(q.steps, s)
}

func (q *Queue) Start() (chan any, error) {
	l := len(q.steps)
	var o chan any
	if l == 0 {
		o = q.output
	} else {
		o = q.steps[0].Input
	}
	go func(output chan any) {
		for i := range q.inputValues {
			output <- i
		}
		close(output)
	}(o)
	for _, step := range q.steps {
		go func(s queueStep) {
			for i := range s.Input {
				v, err := s.ProcessFunc(i)
				if err != nil {
					fmt.Printf("Error while processing input: %v", i)
				} else {
					s.Output <- v
				}
			}
			close(s.Output)
		}(step)
	}
	return q.output, nil
}
