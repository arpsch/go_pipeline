package main

import "fmt"

var version = "dev"

func main() {
	fmt.Println("version - ", version)

	fmt.Println(hello())
}

func hello() string {
	return "Hello Glang"
}
