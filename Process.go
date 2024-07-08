package main

// Controller is a custom type
type Controller struct {
	name string
	id   int
}

// Method with a pointer receiver
func (c *Controller) setNamePointer(name string) {
	c.name = name
}
func (c *Controller) setIdPointer(id int) {
	c.id = id
}

// Method with a value receiver
func (c Controller) setNameValue(name string) {
	c.name = name
}
func (c Controller) SetIdValue(id int) {
	c.id = id
}

// func main() {
// 	c1 := &Controller{name: "Original"}
// 	c2 := Controller{name: "Original"}

// 	// Using the pointer receiver method
// 	c1.setNamePointer("Changed with Pointer")
// 	fmt.Println(c1.name) // Output: Changed with Pointer

// 	// Using the value receiver method
// 	c2.setNameValue("Changed with Value")
// 	fmt.Println(c2.name) // Output: Original
