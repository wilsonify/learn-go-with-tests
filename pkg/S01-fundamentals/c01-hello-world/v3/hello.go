package hello

import "fmt"

// Hello returns a personalised greeting.
func Hello(name string) string {
	return "Hello, " + name
}

func Mainly() {
	fmt.Println(Hello("world"))
}
