package main

import (
	"fmt"
	"solutions/reader"
	"strings"
)

func main() {
	fmt.Print("Enter number: ")
	num := reader.ReadData()

	fmt.Println("all subpalindromes", isPalindromeHard(num))

	if isPalindromeEasy(num) {
		fmt.Printf("Is Palindrome: %v\n", isPalindromeEasy(num))
		return
	}
	fmt.Println("no palindromes was found.")

}

func isPalindromeEasy(s string) bool {
	var b strings.Builder
	b.Grow(len(s))

	for i := len(s) - 1; i >= 0; i-- {
		b.WriteString(string(s[i]))
	}

	if s == b.String() {
		return true
	}
	return false
}

func isPalindromeHard(s string) int {
	var count int
	for i := 0; i < len(s); i++ {
		countSubPalindrome(i, i, &count, s)
	}

	for i := 0; i+1 < len(s); i++ {
		countSubPalindrome(i, i+1, &count, s)
	}

	return count
}

func countSubPalindrome(first int, last int, count *int, s string) {
	for 0 <= first && last < len(s) && s[first] == s[last] {
		*count++
		first--
		last++
	}
}
