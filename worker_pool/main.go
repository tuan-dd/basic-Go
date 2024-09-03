package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type TaskEmail struct {
	email string
	name  string
}

func main() {
	task := []Task{
		&TaskEmail{email: "tuan@gmail.com", name: "tuan"},
		&TaskEmail{email: "tuan1@gmail.com", name: "tuan1"},
		&TaskEmail{email: "tuan2@gmail.com", name: "tuan2"},
		&TaskEmail{email: "tuan3@gmail.com", name: "tuan3"},
		&TaskEmail{email: "tuan4@gmail.com", name: "tuan4"},
		&TaskEmail{email: "tuan5@gmail.com", name: "tuan5"},
		&TaskEmail{email: "tuan6@gmail.com", name: "tuan6"},
		&TaskEmail{email: "tuan7@gmail.com", name: "tuan7"},
		&TaskEmail{email: "tuan8@gmail.com", name: "tuan8"},
		&TaskEmail{email: "tuan9@gmail.com", name: "tuan9"},
	}

	taskRun := WorkerPool{
		Task:        task,
		concurrency: 5,
	}
	taskRun.Run()
}

func (te *TaskEmail) process() {
	fmt.Printf("Prepare send mail %v\n", te.name)

	randomNumber := rand.Intn(10-3+1) + 3
	time.Sleep(time.Second * time.Duration(randomNumber))
	fmt.Printf("Send email successful for %v\n", te.email)
}
