package boolean

import "fmt"

func ExampleOr() {
  fmt.Println(Or(true, false))                 // true
  fmt.Println(Or(false, false))                // false
  fmt.Println(Or4(false, true, false, true))   // true
  fmt.Println(Or4(false, false, false, false)) // false
  // Output:
  // true
  // false
  // true
  // false
}
