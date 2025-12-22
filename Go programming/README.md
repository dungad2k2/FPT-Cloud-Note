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