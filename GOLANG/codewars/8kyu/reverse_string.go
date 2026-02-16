package main

func ReverseString(word string) string {
  runes := []rune(word)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

// OTHER SOLUTIONS SEE
// func Solution(word string) (reversed string) {
  
//   for i := len(word)-1; i >= 0; i-- {
//     reversed += word[i:i+1]
//   }
  
//   return
// }