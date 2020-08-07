package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Here is a simple printing.")
	fmt.Fprintln(os.Stdout, "It implemented by writers.")
	io.WriteString(os.Stdout, "Writer are comes from the io library.")
}
