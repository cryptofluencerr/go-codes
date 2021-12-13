package main

import (
	"fmt"
	// "time"
)

// func pinger(c chan string) {
//   for i := 0; ; i++ {
//     c <- "ping"
//   }
// }

// func printer(c chan string) {
//   for {
//     msg := <- c
//     fmt.Println(msg)
//     time.Sleep(time.Second * 1)
//   }
// }

func main() {
	// var c chan string = make(chan string)

	// go pinger(c)
	// go printer(c)

	// var input string
	// fmt.Scanln(&input)
	var j int = 27
	i := 42
	fmt.Printf("%v, %T", i, i)
  fmt.Print("\n")
	fmt.Printf("%v, %T", j, j)
}
