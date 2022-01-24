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
		description: StringPrompt("Ingrese el nombre del proyecto"),
		status:      StringPrompt("Ingrese el estado del proyecto"),
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
