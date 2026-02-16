package kata

import (
  "fmt"
  "time"
)

func UnluckyDays(year int) (ret int) {
  
  for i := 1; i <= 12; i++ {
    date := time.Date(year, time.Month(i), 13, 0, 0, 0, 0, time.UTC)
    day := date.Weekday()
    if day == time.Friday {
      ret++
    }
    fmt.Println(date)
  }
  return ret
}

// other solutions
// package kata 

// import "time"

// func UnluckyDays(year int) int {
//   tot := 0
//   for m:=1; m<=12; m++ {
//     if time.Date(year, time.Month(m), 13, 0, 0, 0, 0, time.UTC).Weekday() == 5 {
//       tot++
//     }
//   }
// 	return tot
// }