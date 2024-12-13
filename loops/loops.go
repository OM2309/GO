package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Infinite loop

	// for {
	// 	fmt.Println("1")
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	for i := range 3 {
		fmt.Println(i)
	}

}
