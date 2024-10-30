package main

import "fmt"

func main(){
    const occupancyLimit = 12

    var occupancyLimit1 uint8
    var occupancyLimit2 int64
    var occupancyLimit3 float32

    // assign our untyped const to an uint8 variable
    occupancyLimit1 = occupancyLimit
    // assign our untyped const to an int64 variable
    occupancyLimit2 = occupancyLimit
    // assign our untyped const to an float32 variable
    occupancyLimit3 = occupancyLimit

    fmt.Println(occupancyLimit1, occupancyLimit2, occupancyLimit3)
}