package main

import "fmt"

func main() {
	var name string = "Hello GO"
	var name2 string
	name2 = "Charuni"
	name2 = "Isuru"
	var number uint16 = 260
	var num = 230
	numEasy := 520
	fmt.Printf("%T", numEasy)
	fmt.Printf("%T", num)
	fmt.Println("Hello world", name, name2, number)
	//all variables have default val
	var defnum uint64
	var bl bool
	fmt.Println(defnum, bl)

}
