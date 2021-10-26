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
    // tanzim.lastName = "Rabbi"
	tanzimPointer := &tanzim
	tanzimPointer.updateName("Rabbi")
    tanzim.print()
}

func (pointerToPerson *person) updateName(newLastName string)  {
	(*pointerToPerson).lastName = newLastName
}

func (p person) print() {
    fmt.Printf("%+v", p)
}