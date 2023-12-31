package hello

// User user type
type User struct {
	ID   int64
	Name string
	Addr *Address
}

// Address address type
type Address struct {
	City   string
	ZIP    int
	LatLng [2]float64
}

var user1 = User{
	ID:   1,
	Name: "Lucas",
}

// Hello writes a welcome string
func Hello() string {
	return "Hello, my user " + user1.Name
}
