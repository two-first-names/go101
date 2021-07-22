package main

import "fmt"

func main() {
	thing := NewThing(5)
	fmt.Println(thing.GetThingy()) // 5
}
