package main

import (
	endpoint "github.com/mohammad-quanit/go-tdd/Server"
)

func main() {
	endpoint.HttpExample()
}

func HelloWorld() string {
	return "Hello World!"
}

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
