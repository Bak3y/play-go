package main

import "fmt"

const engishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return engishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
