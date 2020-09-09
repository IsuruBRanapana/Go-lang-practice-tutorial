package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func FloatToInt(x float64) int {
	y := int(x)
	return y
}
func main() {
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter first float")
	scanner1.Scan()
	input, _ := strconv.ParseFloat(scanner1.Text(), 32)
	y := FloatToInt(input)
	fmt.Printf("The first integer is %d", y)
	fmt.Println("")
	fmt.Println("Enter second float")
	scanner1.Scan()
	input1, _ := strconv.ParseFloat(scanner1.Text(), 32)
	z := FloatToInt(input1)
	fmt.Printf("The second integer is %d", z)
	fmt.Println("")
}
