    
package main

import (
	"fmt"
)

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
 

func main() {
	fmt.Printf("Absolute value is %v\n", Abs(-9))
	 
}