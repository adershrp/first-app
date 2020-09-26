package main

import (
    "log"
)

func main() {
    log.Println("Start")
    panicker()
    log.Println("End")
}
func panicker() {
    log.Println("About to Panic")
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error", err) // eating the error. hence line10 end will print
            panic(err)                // rethrowing the error. this wil not let line10 to execute
        }
    }()
    panic("Something Happened")
    log.Println("Done Panicing")
}
