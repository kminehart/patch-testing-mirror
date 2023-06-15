package main

import "log"

func Add(a, b int) int {
	return a + b
}

func main() {
	log.Println(Add(1, 2))
}
