package main

import "fmt"

func main() {
	fmt.Printf("hello, world\n")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		if i%5 != 0 && i%3 != 0 {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}
