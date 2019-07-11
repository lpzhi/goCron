package main

import "fmt"

func Test(person string) func() string {

	return func() string {
		return (person + " is working")
	}
}

func main() {
	p := Test("js")

	fmt.Println(p)
}
