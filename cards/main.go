package main

import (
	"fmt"
)

type card struct {
	description string
	status      string
	tasks       []task
}

type task struct {
	description string
	//	deadline    time.Time
	status string
}

func main() {
	project := card{
		description: "first project",
		status:      "in progress",
		tasks: []task{
			{"task 1", "in progress"},
			{"task 2", "done"},
		},
	}
	fmt.Println(project)
}
