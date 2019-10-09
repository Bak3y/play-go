package main

import "fmt"

const engishHelloPrefix = "Hello, "

func Hello(name string) string {
	return engishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
