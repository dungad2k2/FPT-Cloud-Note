package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct{
	URL string
	Status string
}
func checkURL(ctx context.Context, url string) string{
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "Error: Timed out! "
		}
		return fmt.Sprintf("Error: %s", err)
	}
	defer resp.Body.Close()
	return resp.Status
}
func worker(jobs <- chan string, results chan <- Result, wg *sync.WaitGroup,){
	defer wg.Done()
	var res Result 
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for job := range jobs{
		res.URL = job
		res.Status = checkURL(ctx, res.URL)
		results <- res
	}
}
func main(){
	var wg  sync.WaitGroup
	urls := []string{
    "https://www.google.com",
    "https://www.github.com",
    "https://go.dev",
    "https://www.stackoverflow.com",
    "https://doesntexist.xyz", // This one should fail/error
	}

	jobs := make(chan string, len(urls))
	results := make(chan Result, len(urls))
	for i := 1; i <= 2; i++{
		wg.Add(1)
		go worker(jobs, results, &wg)
	}
	for _, url := range urls{
		jobs <- url
	}
	close(jobs)
	go func(){
		wg.Wait()
		close(results)
	}()
	for k := range results{
		fmt.Println(k.URL, " ", k.Status)
	}
}