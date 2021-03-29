package boolean

import "fmt"

func ExampleSelect() {
  fmt.Println(Select(0, true, false))               // true
  fmt.Println(Select(1, true, false))               // false
  fmt.Println(Select4(1, true, true, false, false)) // true
  fmt.Println(Select4(2, true, true, false, false)) // false
  // Output:
  // true
  // false
  // true
  // false
}
