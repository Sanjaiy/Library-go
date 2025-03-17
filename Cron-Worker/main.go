package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	_, err = s.NewJob(gocron.DurationJob(1 * time.Second), gocron.NewTask(task), gocron.WithName("task"))
	if err != nil {
		panic(err)
	}

	_, err = s.NewJob(gocron.DurationJob(2 * time.Second), gocron.NewTask(func() {
		fmt.Println("Vanakam da")
	}), gocron.WithName("task 2"))
	if err != nil {
		panic(err)
	}

	s.Start()

	select {
	case <-time.After(10 * time.Second):
	}

	err = s.Shutdown()
	if err != nil {
		panic(err)
	}
}

func task() {
	fmt.Println("Task executed")
}
