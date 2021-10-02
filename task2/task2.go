package main

import (
	"fmt"
	"log"
	"solutions/reader"
	"strconv"
	"strings"
)

// Input example: 5539 1488 0343 6464

func main() {
	fmt.Print("Enter card number - ")
	cnInput := reader.ReadData()
	cardNumber := strings.Split(cnInput, " ")

	validCard, err := validateCardNumber(cardNumber)
	if err != nil {
		log.Println(err)
	}

	if validCard {
		fmt.Print("---- ---- ---- ")
		for i, cn := range cardNumber {
			if i == 3 {
				fmt.Print(cn)
			}
		}
	} else {
		fmt.Println("Invalid card number")
	}

}

func validateCardNumber(cardNumber []string) (bool, error) {
	if len(cardNumber) != 4 {
		return false, nil
	}

	for _, cn := range cardNumber {
		digit := strings.Split(cn, "")
		if len(digit) != 4 {
			return false, nil
		}
	}

	paymentSystem, err := strconv.Atoi(cardNumber[0])
	if err != nil {
		return false, err
	}
	if paymentSystem/1000 == 4 || paymentSystem/1000 == 5 || paymentSystem/1000 == 3 {
		return true, nil
	}

	return false, nil
}
