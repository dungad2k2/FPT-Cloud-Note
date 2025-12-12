package checker
import (
	"net/http"
	"net/http/httptest"
	"testing"
	"context"
	"time"
)

// TestCheck tests the unexported check() function.
// Since we are in the SAME package ('package checker'), we can see it!
func TestCheck(t *testing.T) {
	// 1. Create a Fake Server
	// This server intercepts requests and replies however we want.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// We force the server to reply with "200 OK"
		w.WriteHeader(http.StatusOK)
	}))
	// Close the server when test finishes
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	// 2. Use the Fake Server's URL
	// server.URL will be something like "http://127.0.0.1:54321"
	status := CheckURL(ctx, server.URL)

	// 3. Assert the result
	expected := "200 OK"
	if status != expected {
		t.Errorf("Expected %s, but got %s", expected, status)
	}
}

func TestCheckFail(t *testing.T) {
	// 1. Create a Fake Server that returns an Error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError) // 500 Error
	}))
	defer server.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	status := CheckURL(ctx, server.URL)

	// In Go, the status string for 500 is "500 Internal Server Error"
	expected := "500 Internal Server Error"
	if status != expected {
		t.Errorf("Expected %s, but got %s", expected, status)
	}
}