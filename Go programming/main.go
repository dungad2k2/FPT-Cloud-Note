package main

import (
	"fmt"
	"os"
	"io"
)

type ByteCounter int
func (b *ByteCounter) Write(p []byte)(int, error){
	*b += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	f1, _ := os.Open("a.txt")
	f2 := &c 
	n, _ := io.Copy(f2, f1)
	fmt.Println(n)
	fmt.Println(c)
}
