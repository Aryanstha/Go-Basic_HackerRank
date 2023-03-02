package main

import (
    "fmt"
    "strconv"
    
)
func cB(n int)int{
    c := 0
    for n!=0{
        c += n&1
        n >>=1
    }
    return c
}
func main(){
    var n int
    fmt.Scanln(&n)
    strconv.FormatInt(int64(n), 2)
    fmt.Printf("%d", cB(n))
}
