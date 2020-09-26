package main

import (
    "fmt"
)

func main() {
    number := 8
    // switch with multiple checks in one case
    switch number {
    case 1, 3, 5, 7, 9:
        fmt.Println("Odd Number")
        fmt.Printf("%T\t%v\n", number, number)
    case 2, 4, 6, 8, 10:
        fmt.Println("Even Number")
        fmt.Printf("%T\t%v\n", number, number)
    default:
        fmt.Println("Another Number")
    }
    
    textSample := "Monday"
    switch textSample {
    case "Sunday", "Saturday":
        fmt.Println("Weekend")
    case "Monday", "Tuesday":
        fmt.Println("Weekday")
        
    }
    
    // switch with initializer
    switch i := 1 + 4; i {
    case 1, 3, 5, 7, 9:
        fmt.Println("Odd Number", i)
    case 2, 4, 6, 8, 10:
        fmt.Println("Even Number", i)
    }
    
    // variance of switch without expression, case level expression
    numberValue := 10454
    switch {
    case numberValue%2 != 0:
        fmt.Println("Odd Number", numberValue)
    case numberValue%2 == 0:
        fmt.Println("Even Number", numberValue)
    }
    
    i := 1
    switch {
    case i%2 != 0:
        fmt.Println("Even", i)
        fallthrough // irrespective of condition my next block wil be executed
    case i%2 == 0:
        fmt.Println("Odd", i)
        fallthrough
    default:
        fmt.Println("Junk")
    }
    
    var z interface{} = [...]int{1, 2, 3}
    // type switch
    switch z.(type) {
    case int:
        fmt.Println(z, "is int type")
        break
        fmt.Println("This Line wont execute") // this line wont execute.
    case float64:
        fmt.Println(z, "is float32 type")
    case string:
        fmt.Println(z, "is string type")
    case [3]int:
        fmt.Println(z, "is [3]int type")
    default:
        fmt.Println(z, "is unknown type")
        
    }
}
