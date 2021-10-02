package main

import (
	"fmt"
	"log"
	"solutions/reader"
	"strconv"
	"strings"
)

// Example of input data: 30,-1,-6,90,-6,22,52,123,2,35,6

func main() {
	fmt.Print("Enter data - ")
	strInput := reader.ReadData()
	numberStr := strings.Split(strInput, ",")

	res, err := countPositiveNumber(numberStr)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Quantity positive numbers - %d", res)
}

func countPositiveNumber(numberStr []string) (int, error) {
	count := 0
	for _, ns := range numberStr {
		number, err := strconv.Atoi(ns)
		if err != nil {
			return 0, err
		}

		if number > 0 {
			count++
		}
	}
	return count, nil
}
