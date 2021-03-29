package boolean

import "fmt"

func ExampleNor() {
  fmt.Println(Nor(false, false))                // true
  fmt.Println(Nor(true, false))                 // false
  fmt.Println(Nor4(false, false, false, false)) // true
  fmt.Println(Nor4(false, false, true, false))  // false
  // Output:
  // true
  // false
  // true
  // false
}
