package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	pages := []string{
		"http://google.com",
		"http://facebook.com",
		"http://microsxzcvzxcvzxcoft.com",
	}

	for _, link := range pages {
		go helthcheck(link)
	}
}

func helthcheck(link string) {

	resp, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " is down.")
		os.Exit(1)
	}

	fmt.Println(link, " is up. Status code: ", resp.StatusCode)

}
