/*
LastDigit returns the last decimal digit of a^b, where a and b are
non-negative integers given as (possibly very large) values.

Key points:
- a and b can be very large (you cannot compute a^b directly).
- The result is the last decimal digit of a^b.
- By convention, 0^0 is defined to be 1.
- The input is always valid.

Examples:
- last digit of 9^7 is 9  (since 9^7 = 4 782 969)
- last digit of (2^200)^(2^300), which has over 10^92 digits, is 6
*/


package main

func LastDigit(n1, n2 string) int {
// Take only the last digit of n1.
// Determine its cycle length (depends on the base).
// Compute n2 % cycleLength.
// Find the corresponding last digit.
    if n2 == "0" {
        return 1 
    }
    
    // last digit of base n1
    lastDigit := int(n1[len(n1)-1] - '0')
  
    // get exponent mod 4 (cycle length = 4 for all except 0, 1, 5, 6)
    mod := 0
    for i := 0; i < len(n2); i++ {
        mod = (mod*10 + int(n2[i]-'0')) % 4
    }
    if mod == 0 {
        mod = 4
    }

    // compute (lastDigit^mod) % 10 safely
    result := 1
    for i := 0; i < mod; i++ {
        result = (result * lastDigit) % 10
    }

    return result

}