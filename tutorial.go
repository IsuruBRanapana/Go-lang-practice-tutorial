package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"math" this package to do more mathematical operations
)

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

	//arithmetic operation
	//variables must in same type
	var num1 int32 = 45
	var num2 float32 = 5.6
	answer := float32(num1) / num2
	fmt.Printf("the answer is %.2f \n", answer)

	//condition and boolean operators
	a := 5
	b := 9.8
	val := float64(a) >= b
	fmt.Printf("%t \n ", val)
	val2 := ((true || false) && !true || 6 > 5) && true
	fmt.Printf("%t \n", val2)

	//if-else if
	nametim := "tim"
	fmt.Println("before if")
	if nametim == "mit" {
		fmt.Println("inside if")
	} else if nametim == "tim" {
		fmt.Println("second inside")
	} else {
		fmt.Println("no no")
	}
	fmt.Println("after if")

	//for loop
	forx := 1
	for forx <= 5 {
		fmt.Println(forx)
		forx++
	}
	//for different
	for fory := 6; fory <= 10; fory++ {
		fmt.Println(fory)
	}
	//break and continue with for
	for forz := 0; forz <= 1000; forz++ {
		if forz != 0 && forz%3 == 0 && forz%7 == 0 && forz%9 == 0 {
			fmt.Println(forz)
			continue
		} else if (forz != 0) && (forz%3 == 1) && (forz%7 == 1) && (forz%9 == 1) {
			fmt.Printf("THE NUM %d \n", forz)
			break
		}
	}

	//switch Statement
	ams := 15
	switch ams {
	case 1:
		fmt.Println("one")
	case 2, -2:
		fmt.Println("two or neg two")
	default:
		println("not one or two")
	}

	//arrays
	var arr [5]int
	arr[0] = 100
	arr[4] = 55
	fmt.Println(arr)

	arr2 := [8]int{6, 7, 8}
	fmt.Println(arr2)

	//2D arrays
	arr3 := [6][2]int{{1, 3}, {3, 5}}
	fmt.Println(arr3[1][1])

	//slices
	var arrs [5]int = [5]int{1, 2, 3, 4, 5}
	var sl []int = arrs[1:3]
	fmt.Println(sl)
	fmt.Println(sl[1:])

	var arrslice []int = []int{4, 8, 9, 7, 2}
	fmt.Println(cap(arrslice[:3]))
	arrslice = append(arrslice, 45)
	fmt.Println(arrslice)

	//range and slice
	var abc []int = []int{1, 2, 3, 4, 5, 6, 9, 8, 7, 4, 56}
	for i := 0; i < len(abc); i++ {
		fmt.Println(abc[i])
	}
	for i, element := range abc {
		fmt.Printf("%d : %d \n", i, element)
	}

	for _, element := range abc {
		fmt.Printf("%d \n", element)
	}
	//avoid duplicate
	for i, element := range abc {
		for j, element2 := range abc {
			if element == element2 && i < j {
				fmt.Println(element)
			}
		}
		fmt.Println("abcda")
		for j := i + 1; j < len(abc); j++ {
			element2 := abc[j]
			if element2 == element {
				fmt.Println(element)
			}
		}
	}

	//commented because to run below one
	//console input
	//scanner := bufio.NewScanner(os.Stdin)
	//fmt.Printf("Type Something : ")
	//scanner.Scan()
	//input := scanner.Text()
	//fmt.Printf("You Typed : %q",input)

	//covert input string in to int or something
	scanner1 := bufio.NewScanner(os.Stdin)
	fmt.Printf("Input Birth year")
	scanner1.Scan()
	newInput, _ := strconv.ParseInt(scanner1.Text(), 10, 64)
	fmt.Printf("You will be %d years old at end of 2020", 2020-newInput)

}
