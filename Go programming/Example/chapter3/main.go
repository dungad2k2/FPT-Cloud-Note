package main

// func printNumber() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func main(){
// 	go printNumber()
// 	fmt.Println("Go routine started")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Main function ended")
// }

// Channel

// func sendData(ch chan int) {
// 	ch <- 10
// }
// func main() {
// 	ch := make(chan int) // Create a channel
// 	go sendData(ch)
// 	data := <-ch
// 	fmt.Println("Receive:", data)
// }

// Buffered Channel

// func main(){
// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// Channel Direction

// func sendOnly(ch chan<- int) {
// 	ch <- 10
// }

// func receiveOnly(ch <-chan int) {
// 	data := <-ch
// 	fmt.Println("Receive:", data)
// }

// func main() {
// 	ch := make(chan int)
// 	go sendOnly(ch)
// 	go receiveOnly(ch)
// 	fmt.Scanln()
// }

// Select Statement

// func main(){
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func(){
// 		time.Sleep(1 * time.Second)
//         ch1 <- "Channel 1"
// 	}()
// 	go func(){
// 		time.Sleep(3 * time.Second)
// 		ch2 <- "Channel 2"
// 	}()
// 	for i := 0; i < 2; i++ {
// 		select {
// 		case msg1 := <-ch1:
// 			fmt.Println("Received", msg1)
// 		case msg2 := <-ch2:
// 			fmt.Println("Received", msg2)
// 		}
// 	}
// }

// WaitGroup Ensuring All Goroutines Finish

// func printNumber(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("Routine %d: %d\n", id, i)
// 	}
// }

// func main(){
// 	var wg sync.WaitGroup
// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go printNumber(i, &wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All routines finished")
// }

// Mutex Used to handle shared data between goroutines

// func main() {
// 	var (
// 		count int
// 		mu sync.Mutex
// 		wg sync.WaitGroup
// 	)
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			mu.Lock()
// 			count++
// 			mu.Unlock()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Count:", count)
// }
