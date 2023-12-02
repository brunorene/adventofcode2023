package main

import (
	"log"
)

func main() {
	day1Part1()
	day1Part2()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
