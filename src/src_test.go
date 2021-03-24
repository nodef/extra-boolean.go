package src

import (
	"fmt"
)

func ExampleParse_true1() {
	fmt.Println(Parse("1"))
	// Output: true
}

func ExampleParse_true2() {
	fmt.Println(Parse("truthy"))
	// Output: true
}

func ExampleParse_true3() {
	fmt.Println(Parse("not off"))
	// Output: true
}

func ExampleParse_false1() {
	fmt.Println(Parse("not true"))
	// Output: false
}

func ExampleParse_false2() {
	fmt.Println(Parse("inactive"))
	// Output: false
}

func ExampleParse_false3() {
	fmt.Println(Parse("disabled"))
	// Output: false
}

// func TestParse(t *testing.T) {
// 	var a, b bool
// 	a = Parse("1")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("1") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("truthy")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("truthy") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("not off")
// 	b = true
// 	if a != b {
// 		t.Errorf(`Parse("not off") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("not true")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("not true") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("inactive")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("inactive") = %t, expected %t`, a, b)
// 	}

// 	a = Parse("disabled")
// 	b = false
// 	if a != b {
// 		t.Errorf(`Parse("disabled") = %t, expected %t`, a, b)
// 	}
// }
