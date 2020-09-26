package main

import (
    "fmt"
)

func main() {
    employeeMap := map[int]string{
        1: "Adersh",
        2: "Divya",
        3: "Prayag",
        4: "Kaira",
    }
    // check and print
    if pop, ok := employeeMap[1]; ok {
        fmt.Println("Employee Name:", pop)
    }
    // check and print
    if pop, ok := employeeMap[8]; ok {
        fmt.Println("Employee Name:", pop)
    }
}
