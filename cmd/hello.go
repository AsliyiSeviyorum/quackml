package main

import "fmt"

const greetingPrefix = "hi, "

func Hello(name string) string {
	if name == "" {
		name = "SUPRA"
	}

	return greetingPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
