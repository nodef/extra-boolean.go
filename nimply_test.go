package boolean

import "fmt"

func ExampleNimply() {
  fmt.Println(Nimply(true, false))  // true
  fmt.Println(Nimply(true, true))   // false
  fmt.Println(Nimply(false, true))  // false
  fmt.Println(Nimply(false, false)) // false
  // Output:
  // true
  // false
  // false
  // false
}
