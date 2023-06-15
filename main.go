package main

import "log"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	log.Println(Add(1, 2))
	log.Println(Sub(2, 1))
}
