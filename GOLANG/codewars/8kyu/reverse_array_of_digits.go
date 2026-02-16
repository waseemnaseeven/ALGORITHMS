package main

/*
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

Example (Input => Output):
35231 => [1,3,2,5,3]
0     => [0]
*/

func Digitize(n int) (ret []int) {
  if n == 0 {
    ret = append(ret, 0)
    return ret
  }
  digit := 0
  for n != 0 {
    digit = n % 10
    ret = append(ret, digit)
    n = n / 10
  }
  return ret
}

// OTHERS
// func Digitize(n int) []int {
// 	var ret []int
// 	for {
// 		ret = append(ret, n%10)
// 		n /= 10
// 		if n == 0 {
// 			return ret
// 		}
// 	}
// }