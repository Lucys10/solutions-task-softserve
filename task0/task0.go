package main

import (
	"fmt"
	"log"
	"solutions/reader"
	"strconv"
)

func main() {
	fmt.Print("Enter height - ")
	h, err := strconv.Atoi(reader.ReadData())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter wight - ")
	w, err := strconv.Atoi(reader.ReadData())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Enter symbol to print - ")
	s := reader.ReadData()

	for i := 0; i < h; i++ {
		if i%2 != 0 {
			fmt.Print("\n" + " ")
		}

		for j := 0; j < w; j++ {
			fmt.Printf("%s ", s)
		}

		if i%2 != 0 {
			fmt.Print("\n")
		}
	}

}
