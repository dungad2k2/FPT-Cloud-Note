# Go class 17: Go does OPP

- The essential elements of object oriented programming have been:
    - abstraction
    - encapsulation
    - polymorphism
    - inheritance
1. Abstraction: decoupling behavior from the implementation details 
2. Encapsulation: hiding implementation details from misuse.
3. Polymorphism literally means "many shapes" - multitypes behind a single interface. Three main types are recognized: ad-hoc, parametric (generic programming), subtype (subclass).
4. Inheritance has conflicting meanings: substitution(subtype) polymorphism, structural sharing of implementation details
   
- Go offers 4 main supports for OOP:
    - encapsulation using the package for visibility control.
    - abstraction and polymorphism using interfaces types.
    - enhanced composition to provide structure sharing.

- Go allows defining methods on any user-defined type, rather than only a "class".

# Go class 18: Methods and Interfaces

- An **interface** specifies adstract behavior in terms of **methods**
   ```
    type Stringer interface{
        String() string
    }
   ```
- A method is a special type of function. It has a receiver parameter before the function name parameter
   ```
    type IntSlice []int
    func (is IntSlice) String() string{
        strs []string
        for _, v := range is {
            strs = append(strs, strconv.Itoa(v))
        }
        return "[" + strings.Join(strs, ";") + "]"
    }
   ```


- An interfaces specifies required behavior as a method set

    ```
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
    ```
- Interface composition:
  - io.ReadWriter is actually defined by Go as two interfaces 
    ```
    type Reader interface {
        Read(p []byte) (n int, err error)
    }
    type Writer interface {
        Write(p []byte) (n int, err error)
    }
    type ReadWriter interface {
        Reader
        Writer
    }
    ```
# Go class 19: Composition
- The fields and methods of an embedded struct are promoted to the level of the embedding structure.
```
type Pair struct {
    Path string
    Hash string
}
type PairWithLength struct{
    Pair
    Length int
}

pl := PairWithLength{Pair{"/usr", "0xfdfe"}, 123}
fmt.Println(pl.Path, pl.Length)
```
- A struct can embed a pointer to another type, promotion of its fields and methods works the same way.
```
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
```
# Go class 20: Interfaces and Methods in Detail 

- Nil interfaces: An interface variable is nil until initialized.
- It really has two parts:
  - a value or pointer of some type
  - a pointer to type information so the correct actual method can be identified. 