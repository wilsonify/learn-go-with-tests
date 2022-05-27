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
	if ok {
		fmt.Sprintf("%s:%d", fileName, fileLine)
	}
	fileName_base := path.Base(fileName)
	fmt.Sprintf("fileName_base = %s", fileName_base)
	fileName_clean := path.Clean(fileName)
	fmt.Sprintf("fileName_clean = %s", fileName_clean)
	fileName_dir := path.Dir(fileName)
	fmt.Sprintf("fileName_dir = %s", fileName_dir)
	fileName_ext := path.Ext(fileName)
	fmt.Sprintf("fileName_ext = %s", fileName_ext)
	fileName_isabs := path.IsAbs(fileName)
	fmt.Sprintf("fileName_isabs = %bool", fileName_isabs)
	fileName_join := path.Join(fileName)
	fmt.Sprintf("fileName_join = %s", fileName_join)
	fileName_head, fileName_tail := path.Split(fileName)
	fmt.Sprintf("fileName_head = %s", fileName_head)
	fmt.Sprintf("fileName_tail = %s", fileName_tail)
	return fileName
}

func main() {
	fmt.Println(Hello())

}
