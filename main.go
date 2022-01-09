package main

import "fmt"

type cards []string

func main() {
	card := cards{}
	card = card.add("new world")
	card.show()
}

func (c cards) show() {
	fmt.Println(c)
}

func (c cards) add(text string) cards {
	return append(c, text)
}
