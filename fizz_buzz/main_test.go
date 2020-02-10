package main

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	n := 15
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	actual := fizzBuzz(n)
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v\nGot %v", expected, actual)
	}
}

// Write a program that outputs the string representation of numbers from 1 to n.
// But for multiples of three it should output “Fizz” instead of the number and
// for the multiples of five output “Buzz”.
// For numbers which are multiples of both three and five output “FizzBuzz”.

