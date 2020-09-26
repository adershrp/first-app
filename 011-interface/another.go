package main

import "fmt"

func main() {
    var db1 database = mysql{}
    var db2 database = mongo{}
    fmt.Println(db1.write("hello"))
    fmt.Println(db1.write("go"))
    // fmt.Println(db1.read(0))
    
    fmt.Println(db2.write("hello"))
    fmt.Println(db2.write("nosql"))
    // fmt.Println(db2.read(1))
}
func rxyzead(d database) {
    x := d.read(0)
    fmt.Println(x)
}

type database interface {
    read(index int) string
    write(content string) int
}

type mysql struct {
    data []string
}

func (m mysql) read(index int) string {
    return m.data[index]
}
func (m mysql) write(content string) int {
    m.data = append(m.data, content)
    return len(m.data)
}

type mongo struct {
    data []string
}

func (m mongo) read(index int) string {
    return m.data[index]
}
func (m mongo) write(content string) int {
    m.data = append(m.data, content)
    return len(m.data)
}
