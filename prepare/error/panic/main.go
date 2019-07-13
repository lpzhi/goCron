package main

import "fmt"

func main() {
	defer func() {
		r := recover()
		err, ok := r.(error)
		if r != nil && ok {
			fmt.Println(err.Error())
		}
		fmt.Println("test")
	}()

	panic("test error")
}
