package main

import "fmt"

//if the declaration is in capital letters then it is accessible to all the packages
// var I int = 27

func main() {
	// fmt.Println(I)
	// var n bool
	// m := 2 == 1
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T", m, m)
	// a := 10 		//1010
	// b := 3			//0011
	// fmt.Println(a & b)	//0010
	// fmt.Println(a | b)	//1011
	// fmt.Println(a ^ b)	//1001 (reverse of & if both are same the not counted)
	// fmt.Println(a &^ b)	//0100 (only those are count in which both are zero)
	// 	a := 10
	// 	fmt.Println(a << 2)
	// 	fmt.Println(a >> 2)
	// var n complex128 = complex(5, 6)
	// fmt.Printf("%v, %T", n, n)

	// s := "this is a string"
	// b := []byte(s)
	// fmt.Printf("%v, %T\n", b, b)
	// const (
	// 	_  = iota
	// 	KB = 1 << (10 * iota)
	// 	MB
	// 	GB
	// 	TB
	// 	PB
	// 	ZB
	// 	YB
	// )
	// fileSize := 4000000000.
	// fmt.Println("KB: ", KB)
	// fmt.Println("MB: ", MB)
	// fmt.Println("GB: ", GB)
	// fmt.Println("TB: ", TB, "\n")

	// fmt.Println("%0.4fTB", fileSize/TB)

	//

	// var a = [...]int{1, 2, 3, 4, 5}
	// b := append(a[:1], a[2:]...)
	// fmt.Println(b)
	// fmt.Println(a)
	// fmt.Printf("Cap: %v, Len: %v", cap(b), len(b))

	// c := b[1:]
	// d := b[:len(b)-1]
	// fmt.Println(c)
	// fmt.Println(d)
	statePopulations := map[string]int{
		"California":   1232132131,
		"Texas":        23242432,
		"Florida":      231312313,
		"New York":     31323432,
		"Pennsylvania": 34354645,
	}
	fmt.Println(statePopulations)
}
