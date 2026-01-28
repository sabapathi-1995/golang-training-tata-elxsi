package main

func main() {

	var p1 struct {
		Id     int
		Name   string
		Email  string
		Mobile string
	} = struct {
		Id     int
		Name   string
		Email  string
		Mobile string
	}{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500"}

	PrintPerson(p1)

}

func PrintPerson(p struct {
	Id     int
	Name   string
	Email  string
	Mobile string
}) { // It is a function
	println()
	println("Id:", p.Id)
	println("Name:", p.Name)
	println("Email:", p.Email)
	println("Mobile:", p.Mobile)

	println("-----------\n")
}
