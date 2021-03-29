package boolean

import "fmt"

func ExampleImply() {
  fmt.Println(Imply(true, true))   // true
  fmt.Println(Imply(false, true))  // true
  fmt.Println(Imply(false, false)) // true
  fmt.Println(Imply(true, false))  // false
  // Output:
  // true
  // true
  // true
  // false
}
