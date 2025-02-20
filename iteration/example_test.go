package iteration

import "fmt"

// Add takes two integers and returns the sum of them.
func ExampleRepeat() {
	result := Repeat("p", 7)
	fmt.Println(result)
	// Output: ppppppp
}
