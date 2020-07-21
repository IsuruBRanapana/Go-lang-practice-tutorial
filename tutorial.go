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
	fmt.Printf("%T %v", num, 15)
	fmt.Println("Hello world", name, name2, number)
	//all variables have default val
	var defnum uint64
	var bl bool
	fmt.Println(defnum, bl)

	//integer format
	//get binary
	fmt.Printf("\n Number : %b", 3435)
	fmt.Printf("\n Number : %o", 3435)
	fmt.Printf("\n Number : %d", 3435)
	//hexa in simple letter
	fmt.Printf("\n Number : %x", 3435)
	//hexa in capital letter
	fmt.Printf("\n Number : %X", 3435)
	//padding
	fmt.Printf("\n Number Pad: %09d", 3435)

	//floating points
	fmt.Printf("\n Number : %e", 3435.15456231545132)
	fmt.Printf("\n Number : %f", 3435.15456231545132)
	fmt.Printf("\n Number : %F", 3435.15456231545132)
	fmt.Printf("\n Number : %g", 3435.15456231545132)
	//width and percision
	fmt.Printf("\n Number : %9f", 3435.15456231545132)
	fmt.Printf("\n Number : %.2f", 3435.15456231545132)
	fmt.Printf("\n Number : %9.2f", 3435.15456231545132)
	fmt.Printf("\n Number : %9.f", 3435.15456231545132)
	//padding
	fmt.Printf("\n Number Pad: %09f", 3435.15456231545132)

	//String formatting
	fmt.Printf("\n Name : %s", "Isuru")
	fmt.Printf("\n Name : %q", "Isuru")
	//width and percision
	fmt.Printf("\n Name : %9s is cool", "Isuru")
	fmt.Printf("\n Name : %-9s is cool", "Isuru")

	//Sprint (store print fmt value in variable
	var x string = fmt.Sprintf("\n Number Pad: %09f", 3435.15456231545132)
	fmt.Println(x)

}
