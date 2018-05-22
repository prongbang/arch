# Go Arch

## How to use:
```
$ go get -u github.com/prongbang/arch
```

## Run from `Makefile`:
Run:
```
$ make run
```
Build:
```
$ make build
```
Install:
```
$ make install
```

# API

## Get Book by paged
GET
```
http://localhost:1323/books?page=1&limit=2
```
Response
```json
[
    {
        "id": 1,
        "name": "Book 1",
        "price": 1000
    },
    {
        "id": 2,
        "name": "Book 2",
        "price": 2000
    }
]
```

## Get Book by Id
GET
```
http://localhost:1323/book/10
```
Response
```json
{
    "id": 10,
    "name": "Book 10",
    "price": 10000
}
```

## Create Book
POST
```
http://localhost:1323/book
```
Request
```json
{
    "id": 109,
    "name": "Book 10",
    "price": 10000
}
```
Response
```json
{
    "id": 109,
    "name": "Book 10",
    "price": 10000
}
```

## Create or Update Book
PUT
```
http://localhost:1323/book/100
```
Request
```json
{
    "name": "Book 100",
    "price": 100000
}
```
Response
```json
{
    "id": 100,
    "name": "Book 100",
    "price": 100000
}
```

## Update Book by Id
PATCH
```
http://localhost:1323/book/109
```
Request
```json
{
    "name": "Book 109",
    "price": 100000
}
```
Response
```json
{
    "id": 100,
    "name": "Book 109",
    "price": 100000
}
```

## Delete Book by Id
DELETE
```
http://localhost:1323/book/99
```
Response
```json
{
    "id": 99,
    "name": "Book 99",
    "price": 99000
}
```

----

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

// -------------------------------------------------

type Book2Implement struct {
    name string
} 

func (i *Book2Implement) SetName(newName string) {
    i.name =  i.name + " " + newName + " : book2"
}   

func (i *Book2Implement) GetName() string {
    return i.name + " : book2"
}

func CreateBook2(newName string) *Book2Implement {
    return &Book2Implement{name: newName}
}

func main() {
    b := Create("Hello") // BookImplement
    b.SetName("Golang")

    Print(b)

    b2 := CreateBook2("Hello") // Book2Implement

    Print(b2)
}

func Print(ib IBooker) {
    fmt.Println(ib.GetName())
}
```

Output:
```
Hello Golang
Hello : book2

Program exited.
```
