package main

import "fmt"

type Person struct {
    first_name string
}

func (a Person) name() string {
    return a.first_name
}

type User struct {
    Person
}

func (a User) name() string {
    return a.first_name
}

func main() {
    bob := Person{"Bob"}
    mike := User{}
    mike.first_name = "Mike"

    fmt.Println(bob.name())
    fmt.Println(mike.name())
}
