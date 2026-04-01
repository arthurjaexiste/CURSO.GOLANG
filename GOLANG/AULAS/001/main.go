package main

import "fmt"

func main() {

	fmt.Println("hello world")
}

func init() {
	fmt.Println("teste")
}

// na linguagem go ele sempre vai printar a func init primeiro