package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")

	if err != nil {
		log.Println("err accrued", err)
	}

	fileErr()
	panicIt()
	reportFatal()
}

func fileErr() {
	f, err := os.Create("no-file.log")

	if err != nil {
		log.Println("err accrued again", err)
	}
	defer f.Close()

	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err accrued, document at no-file.log", err)
	}
	defer f2.Close()
}

func reportFatal() {
	defer foo()
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func foo() {
	fmt.Println("Aborted with err")
}

func panicIt() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		panic(err)
	}
}
