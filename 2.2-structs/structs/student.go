package structs

// Student
// DNI,FirstName,LastName,Date
type Student struct {
	DNI       string
	FirstName string
	LastName  string
	Date      string
}

// Methods
func (s Student) Detail() {
	println("DNI:", s.DNI)
	println("First Name:", s.FirstName)
	println("Last Name:", s.LastName)
	println("Date:", s.Date)
}
