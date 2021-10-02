package main

import (
	"fmt"
	"log"
	"solutions/reader"
	"strconv"
)

func main() {
	fmt.Print("Enter fibonacci number - ")
	fibNumber := reader.ReadData()
	fn, err := strconv.Atoi(fibNumber)
	if err != nil {
		log.Println(err)
	}

	fib := fibonacci()
	for i := 0; i < fn; i++ {
		fmt.Print(fib())
		fmt.Print(" ")
	}

}

func fibonacci() func() int {
	firstNumber, secondNumber := 0, 1
	return func() int {
		res := firstNumber
		firstNumber, secondNumber = secondNumber, firstNumber+secondNumber
		return res
	}
}
