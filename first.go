package main
 
import (
    "fmt"
)
 
type Activity interface {
	GetName() string
	GetAge() int
    Walk() string
    Run() string
}

func doSomething(act Activity) {
	fmt.Printf("This is %s\n", act.GetName())
	fmt.Printf("this is age %d \n",act.GetAge())
    fmt.Println(act.Walk())
	fmt.Println(act.Run())
	
    fmt.Println()
}
 
type Human struct {
	name string
	age int
}
 
type Dog struct {
	name string
	age int
}

type Cat struct{
	name string
	age int
}
 
func (h Human) GetName() string {
    return h.name
}

func (h Human) GetAge() int {
	return h.age
}
 
func (h Human) Walk() string {
    return "I'm walking!"
}
 
func (h Human) Run() string {
    return "I'm running!"
}
 
func (d Dog) GetName() string {
    return d.name
}

func (d Dog) GetAge() int {
	return d.age
}
 
func (d Dog) Walk() string {
    return "Whoop! walking!"
}
 
func (d Dog) Run() string {
    return "Whoop! running!"
}

func (c Cat) GetName() string {
	return c.name
}

func (c Cat) GetAge() int {
	return c.age
}

func (c Cat) Walk() string {
	return "Meaw walking"
}
 
func (c Cat) Run() string {
	return "Meaw running"
}

 
func main() {
    var alice Human = Human{name: "Alice",age:10}
	var bob Dog = Dog{name: "Bob",age:5}
	var spice Cat = Cat{name:"Spice",age:3}
 
    doSomething(alice)
	doSomething(bob)
	doSomething(spice)
}