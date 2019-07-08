package main

import (
	"fmt"
	kyu8 "github.com/imskojs/learn_go_lang/00-Toy_problems/codewars/8kyu"
)

func main() {
	var answer interface{}
	answer =
		kyu8.Points([]string{"3:1", "2:2", "1:2"})
	fmt.Println(answer)

}
