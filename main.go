package main

import "fmt"

func main() {
 fmt.Println(addA("test + "))
}

func addA(a string) string {
 return a + "a"
}
