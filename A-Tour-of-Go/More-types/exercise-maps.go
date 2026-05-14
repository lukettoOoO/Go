package main

import (
	"golang.org/x/tour/wc"
	"strings"
	//"fmt"
)

func WordCount(s string) map[string]int {
	string_fields := strings.Fields(s)
	count := make(map[string]int)
	for _, v := range(string_fields) {
		//fmt.Println(v)
		count[v]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
