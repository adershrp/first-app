package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    s1 := newStudent(1, "Adersh", 35)
    s2 := newStudent(2, "Divya", 28)
    students := []student{*s1, *s2}
    marshal, err := json.Marshal(students)
    if nil != err {
        fmt.Println(err)
        return
    }
    a := string(marshal)
    fmt.Println(a)
    bs := []byte(a)
    fmt.Println(bs)
    
    var uStudents []student
    err = json.Unmarshal(bs, &uStudents)
    if nil != err {
        fmt.Println(err)
        return
    }
    for _, v := range uStudents {
        fmt.Println(v.Name, v.Age, v.Id)
    }
    
}

type student struct {
    Id   int
    Name string
    Age  int
}

func newStudent(id int, name string, age int) *student {
    return &student{Id: id, Name: name, Age: age}
}
