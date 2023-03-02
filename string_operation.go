package main

import (
    "fmt"
    "strings"
)

func encryptString(str string) string {
    // Step 1: Trim all spaces at the start and end of the string
    str = strings.TrimSpace(str)
    
    // Step 2: Remove all digits from 0 to 9
    for i := 0; i <= 9; i++ {
        str = strings.ReplaceAll(str, fmt.Sprintf("%d", i), "")
    }
    
    // Step 3: Reverse the string
    reversed := ""
    for i := len(str) - 1; i >= 0; i-- {
        reversed += string(str[i])
    }
    
    // Return the encrypted string
    return reversed
}

func main() {
    message := " 1nput str1ng w1th 123 n0t so secret d1g1ts "
    encryptedMessage := encryptString(message)
    fmt.Println(encryptedMessage)
}

