// package main

// import (
// 	"fmt"
// )

// func calculateSumOfArray(arr []int, ch chan int) {
// 	sum := 0
// 	for _, value := range arr {
// 		sum += value
// 	}
// 	//fmt.Println("Sum of array:", sum)
// 	ch <- sum

// }
// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	ch := make(chan int)
// 	go calculateSumOfArray(arr[:len(arr)/2], ch)
// 	go calculateSumOfArray(arr[len(arr)/2:], ch)
// 	x, y := <-ch, <-ch
// 	fmt.Println("Sum of array:", x+y)
// }