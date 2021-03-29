package boolean

import "fmt"

func ExampleNot() {
  fmt.Println(Not(false)) // true
  fmt.Println(Not(true))  // false
  // Output:
  // true
  // false
}
