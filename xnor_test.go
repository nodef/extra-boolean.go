package boolean

import "fmt"

func ExampleXnor() {
  fmt.Println(Xnor(true, true))                // true
  fmt.Println(Xnor(false, true))               // false
  fmt.Println(Xnor4(true, true, false, false)) // true
  fmt.Println(Xnor4(true, true, true, false))  // false
  // Output:
  // true
  // false
  // true
  // false
}
