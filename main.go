package main

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

func SumAll(numbersTosum ...[]int) []int {
	return nil
}
