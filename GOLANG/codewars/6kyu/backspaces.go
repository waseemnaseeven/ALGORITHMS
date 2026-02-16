/*

Assume "#" is like a backspace in string. This means that string "a#bc#d" actually is "bd"

Your task is to process a string with "#" symbols.

Examples
"abc#d##c"      ==>  "ac"
"abc##d######"  ==>  ""
"#######"       ==>  ""
""              ==>  ""

*/

package kata

func CleanString(s string) string {
  stack := []rune{}

  for _, ch := range s {
    if ch == '#' {
      if len(stack) > 0 {
        stack = stack[:len(stack)-1]
      }
    } else {
        stack = append(stack, ch)
      }
  }
  return string(stack)
}
