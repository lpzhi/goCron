package main

import "fmt"

func lengthOfStr(s string) int {
	lastOcru := make(map[byte]int)
	start := 0
	maxlengh := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOcru[ch]; ok && lastI > start {
			start = i + 1
		}

		if i-start+1 > maxlengh {
			maxlengh = i - start + 1
		}
		lastOcru[ch] = i
	}

	return maxlengh
}

func main() {
	str := "bc12345c67890plkmghb6c"
	i := lengthOfStr(str)
	fmt.Println(i)
}
