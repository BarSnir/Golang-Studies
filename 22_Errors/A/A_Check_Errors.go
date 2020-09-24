package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	exampleB(exampleA())
	exampleC()
}

func exampleA() string {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	answer := strings.Join([]string{answer1, answer2, answer3}, " ")
	return answer
}

func exampleB(output string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}

	dir = strings.Join([]string{dir, `/22_Errors/error.log`}, "")

	fmt.Println("Error logs been stored in ", dir)
	f, err := os.Create(dir)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	r := strings.NewReader(output)

	io.Copy(f, r)
}

func exampleC() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}

	dir = strings.Join([]string{dir, `/22_Errors/error.log`}, "")

	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("And what the file said ", "'", string(bs), "'")
}
