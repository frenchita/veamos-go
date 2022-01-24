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

	lines := 0

	file, err := os.Open("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines++
		quote := scanner.Text()
		password := []byte(quote)

		hashedPassword := hash(password)
		fmt.Println(string(hashedPassword))
	}

	fmt.Println(lines, "hashes have been generated")
	fmt.Printf("time: %s", time.Since(start))

}

func hash(password []byte) []byte {

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return hashedPassword

}
