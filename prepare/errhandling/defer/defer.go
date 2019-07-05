package main

import (
	"fmt"
)

func main() {
	//b := fbi.Fibonacci()
	//
	//fmt.Println(b())
	//fmt.Println(b())
	//fmt.Println(b())
	//fmt.Println(b())
	//fmt.Println(b())
	//fmt.Println(b())

	tryDefer()

}

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
}
