# Go Arch

## How to use:
```
$ go get -u github.com/prongbang/arch
```

Interface and Implements Example:
```go
package main

import "fmt"

type IBooker interface {
    SetName(newName string)
    GetName() string
}

type BookImplement struct {
    name string
} 

func (i *BookImplement) SetName(newName string) {
    i.name =  i.name + " " + newName
}   

func (i *BookImplement) GetName() string {
    return i.name
}

func Create(newName string) *BookImplement {
    return &BookImplement{name: newName}
}

func main() {
    var b IBooker = Create("Hello")
    b.SetName("Golang")
    fmt.Println(b.GetName())
}
```

Output:
```
Hello Golang

Program exited.
```