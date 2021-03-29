package boolean

import "fmt"

func ExampleParse() {
	fmt.Println(Parse("1"))            // true
	fmt.Println(Parse("truthy"))       // true
	fmt.Println(Parse("Not Off"))      // true
	fmt.Println(Parse("Not Inactive")) // true
	fmt.Println(Parse("cold"))         // false
	fmt.Println(Parse("inactive"))     // false
	fmt.Println(Parse("Negative Yes")) // false
	fmt.Println(Parse("Negative Aye")) // false
	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}
