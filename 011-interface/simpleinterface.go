package main

import "fmt"

func main() {
    var h human
    h = vegetarian{
        name: "carrot",
        cal:  50,
    }
    fmt.Println(h.eat("vegetable"))
    
    h = vegetarian{
        name: "beetroot",
        cal:  50,
    }
    fmt.Println(h.eat("vegetable"))
    
    h = nonvegetarian{
        name: "chicken",
        cal:  400,
    }
    fmt.Println(h.eat("nonveg"))
}

// interface define behavior
type human interface {
    eat(string) (int, error)
}

// struct holds data
type vegetarian struct {
    name string
    cal  int
}

// struct holds data
type nonvegetarian struct {
    name string
    cal  int
}

// * method defined for struct has same name as the interface.
// * implicitly implementing interface
func (v vegetarian) eat(food string) (int, error) {
    fmt.Println("Eating", v.name)
    cal := len(food)
    return cal, nil
}

func (n nonvegetarian) eat(food string) (int, error) {
    fmt.Println("Eating", n.name)
    cal := len(food)
    return cal, nil
}
