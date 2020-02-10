// Write a program that outputs the string representation of numbers from 1 to n.
// But for multiples of three it should output “Fizz” instead of the number and
// for the multiples of five output “Buzz”.
// For numbers which are multiples of both three and five output “FizzBuzz”.
package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	result := []string{}

	for i := 1; i <= n; i++ {
		str := ""

		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}

		if len(str) != 0 {
			result = append(result, str)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func main() {
	num := 15
	result := fizzBuzz(num)
	fmt.Println(result)
}
