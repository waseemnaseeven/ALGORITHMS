package main

import (
  "math"
)

/*

If John goes to the cinema 3 times:

System A : 15 * 3 = 45
System B : 500 + 15 * 0.90 + (15 * 0.90) * 0.90 + (15 * 0.90 * 0.90) * 0.90 ( = 536.5849999999999, no rounding for each ticket)
John wants to know how many times he must go to the cinema so that the final result of System B, when rounded up to the next dollar, will be cheaper than System A.

The function movie has 3 parameters: card (price of the card), ticket (normal price of a ticket), perc (fraction of what he paid for the previous ticket) and returns the first n such that

ceil(price of System B) < price of System A.
More examples:
movie(500, 15, 0.9) should return 43 
    (with card the total price is 634, with tickets 645)
movie(100, 10, 0.95) should return 24 
    (with card the total price is 235, with tickets 240)
*/

func Movie(card, ticket int, perc float64) int {
  
  current := float64(ticket) * perc
  sysB := float64(card)
  n := 1.0
  for {
    sysA := float64(ticket) * n
    sysB += current
    current *= perc
    if math.Ceil(sysB) < sysA {
      break
    }
    n++
  }
  return int(n)
}
