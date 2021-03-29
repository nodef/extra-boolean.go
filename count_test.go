package boolean

import "fmt"

func ExampleCount() {
  fmt.Println(Count(true, true))                 // 2
  fmt.Println(Count(true, false))                // 1
  fmt.Println(Count4(true, true, true, false))   // 3
  fmt.Println(Count4(false, true, false, false)) // 1
  // Output:
  // 2
  // 1
  // 3
  // 1
}
