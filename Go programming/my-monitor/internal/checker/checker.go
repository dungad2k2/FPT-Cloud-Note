package checker // The package name matches the folder name

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Result struct{
	URL string
	Status string
}
func CheckURL(ctx context.Context, url string) string{
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
func Worker(jobs <- chan string, results chan <- Result){
	var res Result 
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for job := range jobs{
		res.URL = job
		res.Status = CheckURL(ctx, res.URL)
		results <- res
	}
}