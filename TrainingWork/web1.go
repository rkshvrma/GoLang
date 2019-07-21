package main

 import (
	"fmt"
)

 var strarray = []string{"lorem", "ipsum", "dolor", "sit", "amet"}
 
func main() {

 	for i := 0; i != 5; i++ {

 		fmt.Println("\t", strarray[i])
 
	}
 
}