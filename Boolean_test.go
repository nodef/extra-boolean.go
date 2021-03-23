package Boolean

import (
	"fmt"
	"testing"
)

func ExampleParse() {
	fmt.Println(Parse("true"))
	// Output: true
}

func TestParse(t *testing.T) {
	var a, b bool
	a = Parse("1")
	b = true
	if a != b {
		t.Errorf(`Parse("1") = %t, expected %t`, a, b)
	}

	a = Parse("truthy")
	b = true
	if a != b {
		t.Errorf(`Parse("truthy") = %t, expected %t`, a, b)
	}

	a = Parse("not off")
	b = true
	if a != b {
		t.Errorf(`Parse("not off") = %t, expected %t`, a, b)
	}

	a = Parse("not true")
	b = false
	if a != b {
		t.Errorf(`Parse("not true") = %t, expected %t`, a, b)
	}

	a = Parse("inactive")
	b = false
	if a != b {
		t.Errorf(`Parse("inactive") = %t, expected %t`, a, b)
	}

	a = Parse("disabled")
	b = false
	if a != b {
		t.Errorf(`Parse("disabled") = %t, expected %t`, a, b)
	}
}
