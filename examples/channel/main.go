package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/okieoth/ppipel/examples/queue"
)

type Complex struct {
	S string
	I int
}

func simple(inputSize int) {
	fmt.Println(":)")
	input1 := make(chan int, 100)
	input2 := make(chan string, 100)
	output := make(chan Complex, 100)

	go func(output chan<- int) {
		for i := 0; i < inputSize; i++ {
			sleepTime := rand.Intn(2) + 1
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			output <- i
		}
		close(output)
	}(input1)
	go func(input <-chan int, output chan<- string) {
		for v := range input {
			sleepTime := rand.Intn(20) + 1
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			s := ""
			for i := 0; i < v%10; i++ {
				s += "X"
			}
			output <- fmt.Sprintf("%d: %s", v, s)
		}
		close(output)
	}(input1, input2)
	go func(input <-chan string, output chan<- Complex) {
		currentCount := 0
		for v := range input {
			sleepTime := rand.Intn(200) + 1
			time.Sleep(time.Duration(sleepTime) * time.Millisecond)
			output <- Complex{
				I: currentCount,
				S: v,
			}
			currentCount++
		}
		close(output)
	}(input2, output)
	for o := range output {
		fmt.Printf("Counter: %d, Value: %s\n", o.I, o.S)
	}
	fmt.Printf("Should have %d lines in output\n", inputSize)
}

func example_queue() {
	fmt.Println(":)")

	q := queue.NewQueue(queue.SimpleIterator)
	q.AddStep(func(input any) (any, error) {
		sleepTime := rand.Intn(20) + 1
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		s := ""
		v, ok := input.(int)
		if ok {
			for i := 0; i < v%10; i++ {
				s += "X"
			}
			return fmt.Sprintf("%d: %s", v, s), nil

		} else {
			return 0, fmt.Errorf("Couldn't cast input into int value: input=%v", input)
		}
	})
	currentCount := 0
	q.AddStep(func(input any) (any, error) {
		sleepTime := rand.Intn(200) + 1
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		currentCount++
		v, ok := input.(string)
		if ok {
			return Complex{
				I: currentCount,
				S: v,
			}, nil
		} else {
			var r Complex
			return r, fmt.Errorf("Couldn't cast input into string value: input=%v", input)
		}
	})

	outputChan, _ := q.Start()

	for o := range outputChan {
		v, ok := o.(Complex)
		if ok {
			fmt.Printf("Counter: %d, Value: %s\n", v.I, v.S)
		} else {
			fmt.Printf("Couldn't cast final value to Complex: %v", o)
		}
	}
	fmt.Printf("Should have %d lines in output\n", 100)
}

func main() {
	example_queue()
	simple(100)
}
