package main

import (
	"fmt"
	kyu7 "github.com/imskojs/learn_go_lang/00-Toy_problems/codewars/7kyu"
)

func main() {
	var answer interface{}
	answer = kyu7.GetCount("hello")
	fmt.Println(answer)
}
