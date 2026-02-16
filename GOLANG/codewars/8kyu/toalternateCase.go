package kata

func ToAlternatingCase(str string) (string) {
  runes := []rune(str)
  
  for i := 0; i <= len(runes)-1; i++ {
    if runes[i] >= 'a' && runes[i] <= 'z' {
      runes[i] -= 32
    } else if runes[i] >= 'A' && runes[i] <= 'Z' {
      runes[i] += 32
    }
  }
  
  return string(runes)
}