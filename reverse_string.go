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
    words := strings.Fields(scanner.Text())
    
    if len(words) < 3 {
        fmt.Println("String kurang dari 3 kata)
        return
    }
    
    for i := len(words)-1; i >= 0; i-- {
        fmt.Print(reverseString(words[i]), " ")
    }
}
