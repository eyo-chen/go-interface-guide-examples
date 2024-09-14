package main

// Car is an interface with a method Drive()
type Car interface {
	Drive()
}

// Tesla is a struct that implements the Car interface
// It has a method Drive() that prints "driving a tesla"
type Tesla struct{}

func (t Tesla) Drive() {
	println("driving a tesla")
}

// DriveCar is a function that takes a Car and calls its Drive() method
func DriveCar(c Car) {
	c.Drive()
}

func main() {
	t := Tesla{}
	DriveCar(t)

	/*
		Output:
		driving a tesla
	*/
}
