package main

import "fmt"

func main() {
 fmt.Println("start")
 fmt.Println(addA("test + "))
 fmt.Println("finish")
}

func addA(a string) string {
 return a + "a"
}
