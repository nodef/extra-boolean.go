package boolean

import "fmt"

func ExampleEq() {
  fmt.Println(Eq(true, true))   // true
  fmt.Println(Eq(false, false)) // true
  fmt.Println(Eq(true, false))  // false
  fmt.Println(Eq(false, true))  // false
  // Output:
  // true
  // true
  // false
  // false
}
