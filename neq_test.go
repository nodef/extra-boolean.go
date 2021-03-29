package boolean

import "fmt"

func ExampleNeq() {
  fmt.Println(Neq(true, false))  // true
  fmt.Println(Neq(false, true))  // true
  fmt.Println(Neq(true, true))   // false
  fmt.Println(Neq(false, false)) // false
  // Output:
  // true
  // true
  // false
  // false
}
