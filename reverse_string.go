package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	
	words := strings.Fields(input)
	
	if len(words) < 3 {
		fmt.Println("String kurang dari 3 kata!")
		return
	}
	
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}
	
	for i := 0; i < len(words)/2; i++ {
		j := len(words) - 1 - i
		words[i], words[j] = words[j], words[i]
	}
	
	result := strings.Join(words, " ")
	fmt.Println(result)
}
