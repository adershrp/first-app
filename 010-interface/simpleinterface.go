package main

import (
    "fmt"
)

func main() {
    polo := car{
        model: "polo", brand: "vw", wheels: 4,
    }
    express := train{
        bogies: 18, isElectric: true,
    }
    fmt.Println(polo.drive())
    fmt.Println(express.drive())
    bar(polo)
    bar(express)
    
}

type vehicle interface {
    drive() string
}

func bar(v vehicle) {
    switch v.(type) {
    case car:
        fmt.Println("Car ", v.(car).wheels)
    case train:
        fmt.Println("Train", v.(train).bogies)
    }
}

type car struct {
    model, brand string
    wheels       int
}

func (c car) drive() string {
    return fmt.Sprint(c.brand, " ", c.model, " has ", c.wheels, " wheels.")
}

type train struct {
    isElectric bool
    bogies     int
}

func (t train) drive() string {
    return fmt.Sprint("Train has ", t.bogies, " bogies")
}
