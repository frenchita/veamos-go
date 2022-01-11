package main

import "fmt"

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

func (c cards) add() cards {

	project := card{
		description: "first project",
		status:      "in progress",
		tasks: []task{
			{"task 1", "in progress"},
			{"task 2", "done"},
		},
	}

	c = append(c, project)
	return c

}

func (c cards) show() {

	fmt.Println(c)

}
