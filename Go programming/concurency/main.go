// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )
// // func worker(id int, wg *sync.WaitGroup){
// // 	defer wg.Done()
// // 	fmt.Printf("Worker %d starting\n", id)
// // 	time.Sleep(time.Second)
// // 	fmt.Printf("Worker %d done \n", id)
// // }
// func worker(id int, jobs <-chan int, results chan<- int){
// 	for job := range jobs{
// 		fmt.Printf("Worker %d start job %d\n", id, job)
// 		time.Sleep(time.Second)
// 		fmt.Printf("Worker %d finish job %d\n", id, job)
// 		results <- job * 2 
// 	}
// }
// type SafeCounter struct {
// 	mu sync.Mutex
// 	value int
// }
// func (c *SafeCounter) Increment(){
// 	c.mu.Lock()
// 	c.value++
// 	//c.mu.Unlock()
// }
// func main(){
// 	// var wg sync.WaitGroup
// 	// counter := SafeCounter{}
// 	// for i := 1; i <= 1000; i++{
// 	// 	wg.Add(1)
// 	// 	go func(){
// 	// 		defer wg.Done()
// 	// 		counter.Increment()
// 	// 	}()
// 	// }
// 	// //fmt.Println("Main is waiting for workers.....")
// 	// wg.Wait()
// 	// fmt.Println(counter.value)
// 	// fmt.Println("All workers completed!")
// 	const numJobs = 5
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
// 	for w := 1; w <= 3; w++{
// 		go worker(w, jobs, results)
// 	}
// 	for j := 1; j <= numJobs; j++{
// 		jobs <- j
// 	}
// 	close(jobs)
// 	for a := 1; a <= numJobs; a++ {
// 		<-results
// 	}
// 	// fmt.Println("Main is done sending jobs. Waiting 2 seconds...")
//     // time.Sleep(2 * time.Second) // FORCE main to stay alive
//     // fmt.Println("Main exiting now.")
// }
package main

import (
    "fmt"
  //  "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    // go func() {
    //     time.Sleep(1 * time.Second)
    //     c1 <- "one"
	// //	fmt.Println("first")
    // }()
    // go func() {
    //     time.Sleep(2 * time.Second)
    //     c2 <- "two"
	// 	//fmt.Println("second")
    // }()
	// go func(){
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "c2"
	// }()

    for range 1 {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}