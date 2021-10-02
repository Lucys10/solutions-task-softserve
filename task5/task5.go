package main

import (
	"fmt"
	"log"
	"solutions/reader"
	"strconv"
)

func main() {
	fmt.Print("Enter number Min - ")
	dmn := reader.ReadData()
	numMin, err := strconv.Atoi(dmn)
	if err != nil {
		log.Println(err)
	}

	fmt.Print("Enter number Max - ")
	dmx := reader.ReadData()
	numMax, err := strconv.Atoi(dmx)
	if err != nil {
		log.Println(err)
	}

	simpleCountLuckyTicket := simpleCountTicket(numMin, numMax)
	hardCountLuckyTicker := hardCountTicket(numMin, numMax)

	fmt.Println("--Result--")
	fmt.Printf("EasyFormula: %d\n", simpleCountLuckyTicket)
	fmt.Printf("HardFormula: %d\n", hardCountLuckyTicker)

	if simpleCountLuckyTicket > hardCountLuckyTicker {
		fmt.Println("In a given interval counting lucky tickets will give - EasyFormula")
		return
	}
	fmt.Println("In a given interval counting lucky tickets will give - HardFormula")
}

func simpleCountTicket(numMin, numMax int) int {
	countLuckyTicket := 0
	for i := numMin; i <= numMax; i++ {
		numLeft := i / 1000
		numRight := i % 1000

		numLeft1 := numLeft / 100
		numLeft2 := numLeft / 10 % 10
		numLeft3 := numLeft % 10

		numRight1 := numRight / 100
		numRight2 := numRight / 10 % 10
		numRight3 := numRight % 10

		sumLeft := numLeft1 + numLeft2 + numLeft3
		sumRight := numRight1 + numRight2 + numRight3

		if sumLeft == sumRight {
			countLuckyTicket++
		}
	}
	return countLuckyTicket
}

func hardCountTicket(numMin, numMax int) int {
	countLuckyTicket := 0

	for i := numMin; i <= numMax; i++ {
		numbers := make([]int, 0, 6)
		numLeft := i / 1000
		numRight := i % 1000

		num1 := numLeft / 100
		num2 := numLeft / 10 % 10
		num3 := numLeft % 10

		num4 := numRight / 100
		num5 := numRight / 10 % 10
		num6 := numRight % 10

		numbers = append(numbers, num1, num2, num3, num4, num5, num6)

		sumEven, sumOld := evenAndOld(numbers)

		if sumOld == sumEven {
			countLuckyTicket++
		}
	}
	return countLuckyTicket
}

func evenAndOld(numbers []int) (sumEven, sumOld int) {
	for _, n := range numbers {
		if n%2 == 0 {
			sumEven += n
		} else {
			sumOld += n
		}
	}
	return
}
