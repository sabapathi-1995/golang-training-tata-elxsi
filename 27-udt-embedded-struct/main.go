package main

func main() {

	p1 := Person{Id: 101, Name: "Jiten", Address: struct { // embedded structure
		Line1   string
		City    string
		Pincode string
	}{Line1: "PRRA67", City: "Trivandrum", Pincode: "690511"}}

	p1.PrintPerson()

	//p2 := new(Person)

	//var p2 Person

	// p2 := Person{}

	p2 := &Person{}

	p2.Id = 101
	p2.Name = "Priya"
	p2.Address.Line1 = "HA101"
	p2.Address.City = "Chennai"
	p2.Address.Pincode = "490011"
	p2.PrintPerson()
}

type Person struct {
	Id      int
	Name    string
	Address struct { // embedded structure
		Line1   string
		City    string
		Pincode string
	}
}

func (p *Person) PrintPerson() { // It is a function
	println()
	println("Id:", p.Id)
	println("Name:", p.Name)
	println("Address:")

	println("Line1:", p.Address.Line1)
	println("City:", p.Address.City)
	println("Pincode:", p.Address.Pincode)

	println("-----------\n")
}
