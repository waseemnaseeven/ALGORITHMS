package main

import "math"

/*
Write a function which takes a number as input and returns the sum of the absolute value of each of the number's decimal digits.

For example: (Input --> Output)

10 --> 1
99 --> 18
-32 --> 5
Let's assume that all numbers in the input will be integer values.

*/

func SumDigits(number int) int {
    
    digit := 0
    sum := 0
    for number != 0 {
      digit = number % 10
      sum += digit
      number = number / 10
    }
  ret := math.Abs(float64(sum)) // ou if n < 0 {n=-n}
  return int(ret)
}