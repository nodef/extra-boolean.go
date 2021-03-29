package boolean

import "fmt"

func ExampleEqv() {
  fmt.Println(Eqv(true, true))   // true
  fmt.Println(Eqv(false, false)) // true
  fmt.Println(Eqv(true, false))  // false
  fmt.Println(Eqv(false, true))  // false
  // Output:
  // true
  // true
  // false
  // false
}
