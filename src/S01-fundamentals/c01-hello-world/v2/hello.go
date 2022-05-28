package main

import (
	"fmt"
	"path"
	"runtime"
)

// Hello returns a greeting.
func Hello() string {
	return "Hello, world"
}

func here() string {
	_, fileName, fileLine, ok := runtime.Caller(1)
	fmt.Printf("fileName = %s \n", fileName)
	fmt.Printf("fileLine = %d \n", fileLine)
	fmt.Printf("ok = %t \n", ok)
	fileName_base := path.Base(fileName)
	fmt.Printf("fileName_base = %s \n", fileName_base)
	fileName_clean := path.Clean(fileName)
	fmt.Printf("fileName_clean = %s", fileName_clean)
	fileName_dir := path.Dir(fileName)
	fmt.Printf("fileName_dir = %s", fileName_dir)
	fileName_ext := path.Ext(fileName)
	fmt.Printf("fileName_ext = %s", fileName_ext)
	fileName_isabs := path.IsAbs(fileName)
	fmt.Printf("fileName_isabs = %t", fileName_isabs)
	fileName_join := path.Join(fileName)
	fmt.Printf("fileName_join = %s", fileName_join)
	fileName_head, fileName_tail := path.Split(fileName)
	fmt.Printf("fileName_head = %s", fileName_head)
	fmt.Printf("fileName_tail = %s", fileName_tail)

	return fileName_head
}

func main() {
	fmt.Println(Hello())
	fmt.Println(here())
}
