package main

import (
	"fmt"
	"math"
)

func fizzBuzz () string {
	for i := 0; i < 100; i++ {
		if math.Mod(float64(i), 3) == 0 && math.Mod(float64(i), 5) == 0 {
			fmt.Println("FizzBuzz", i)
			} else if math.Mod(float64(i), 3) == 0 {
			 fmt.Println("Fizz") 
			} else if math.Mod(float64(i), 5) == 0 {
				fmt.Println("Buzz")
			} else {
				fmt.Println(i)
			}
		}
		return "done"
}

func main() {
	fizzBuzz()
}