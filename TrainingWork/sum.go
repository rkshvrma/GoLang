package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter digits ")
	str, err := getline()
	if err == nil {
		fmt.Println("Here is the res:", sumnums(str))
	}
}

func getline() (string, error) {
	return bufio.NewReader(os.Stdin).ReadString('\n')
}

func sumnums(str string) float64 {
	var sum float64
	for _, v := range strings.Fields(str) {
		i, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Println(err)
		} else {
			sum += i
		}
	}
	return sum
}