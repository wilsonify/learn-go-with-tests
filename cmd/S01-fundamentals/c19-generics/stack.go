package main

import (
	"fmt"

	generics "learn.go/S01-fundamentals/c19-generics/v1"
)

func main() {
	fmt.Println("make a new stack")
	myStackOfInts := new(generics.Stack[int])
	fmt.Println("check stack is empty")
	fmt.Println("myStackOfInts.IsEmpty =", myStackOfInts.IsEmpty())

	fmt.Println("add a thing, then check it's not empty")
	fmt.Println("push 123 to stack")
	myStackOfInts.Push(123)
	fmt.Println("myStackOfInts.IsEmpty =", myStackOfInts.IsEmpty())

	fmt.Println("add another thing, pop it back again")
	fmt.Println("push 456 to stack")
	myStackOfInts.Push(456)
	fmt.Println("Pop!")
	value, _ := myStackOfInts.Pop()
	fmt.Println("value =", value, "expected = ", 456)
	fmt.Println("Pop!")
	value, _ = myStackOfInts.Pop()
	fmt.Println("value =", value, "expected = ", 123)
	fmt.Println("myStackOfInts.IsEmpty =", myStackOfInts.IsEmpty())

	fmt.Println("can get the numbers we put in as numbers, not untyped interface{}")
	fmt.Println("Push a 1")
	myStackOfInts.Push(1)
	fmt.Println("Push a 2")
	myStackOfInts.Push(2)
	fmt.Println("Pop!")
	firstNum, _ := myStackOfInts.Pop()
	fmt.Println("Pop!")
	secondNum, _ := myStackOfInts.Pop()
	fmt.Println("firstNum+secondNum =", firstNum+secondNum, "expected = ", 3)

}
