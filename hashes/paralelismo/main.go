package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	start := time.Now()

	c := make(chan string)
	lines := 0

	file, err := os.Open("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
		quote := scanner.Text()
		password := []byte(quote)

		go hash(password, c)

	}

	for i := 0; i < lines; i++ {
		fmt.Println(<-c)
	}

	fmt.Println(lines, "hashes have been generated")
	fmt.Printf("time: %s", time.Since(start))

}

func hash(password []byte, c chan string) {

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	c <- string(hashedPassword)

}
