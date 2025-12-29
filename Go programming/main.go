package main

import (
	"fmt"
	"sort"
	// "os"
	// "io"
)

type ByteCounter int
func (b *ByteCounter) Write(p []byte)(int, error){
	*b += ByteCounter(len(p))
	return len(p), nil
}
type Pair struct {
	Path string 
	Hash string
}
type PairWithLength struct{
	Pair 
	Length int 
}
type Organ struct {
	Name string
	Weight int
}
type Organs []Organ
func (s Organs) Len() int { return len(s)}
func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
} 
type ByName struct {
	Organs
}
type ByWeight struct {
	Organs 
}
func (s ByName) Less (i, j int) bool {
	return s.Organs[i].Name < s.Organs[j].Name
}
func (s ByWeight) Less (i, j int) bool {
	return s.Organs[i].Weight < s.Organs[j].Weight
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}}
	sort.Sort(ByWeight{s})
	fmt.Println(s)
	sort.Sort(ByName{s})
	fmt.Println(s)
}
