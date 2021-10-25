package main

import "fmt"


type contactInfo struct {
    email string
    zipCode int
}


type person struct {
    firstName string
    lastName string
    contactInfo
}

func main() {
    tanzim := person{
        firstName: "Tanzim", 
        lastName: "Rizwan",
        contactInfo: contactInfo{
            email: "tanzim@mail.com",
            zipCode: 1219,
        },
    }
    tanzim.lastName = "Rabbi"
    tanzim.print()
}

func (p person) print() {
    fmt.Printf("%+v", p)
}