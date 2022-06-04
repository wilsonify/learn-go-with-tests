package hello

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello returns a personalised greeting, defaulting to Hello, world if an empty name is passed.
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func Mainly() {
	fmt.Println(Hello("world"))
}
