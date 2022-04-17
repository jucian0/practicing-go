package main

import "fmt"

func for_loop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
}

func while_loop() {
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}

func foreach_loop() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, value := range list {
		fmt.Println(i, value)
	}
}

func exit_loop() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}

func main() {
	//for_loop()
	//while_loop()
	//foreach_loop()
	exit_loop()
}
