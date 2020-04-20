// https://www.hackerrank.com/challenges/string-construction/problem

package main

import "fmt"

func stringConstruction(s string) int32 {
	var exists [26]bool
	var count = int32(0)

	for _, chr := range s {
		exists[chr - 'a'] = true
	}

	for _, exist := range exists {
		if exist {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(stringConstruction("abcd"))
	fmt.Println(stringConstruction("abab"))
}
