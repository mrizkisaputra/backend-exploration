package basic_test

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int32
}

// factory function
func NewPerson(name string, age int32) *Person {
	return &Person{name, age}
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hello, %s your age now %d", p.Name, p.Age)
}

func (p *Person) HappyBirthday() {
	p.Age += 1
}

func TestSturct(t *testing.T) {
	//person := NewPerson("Alice", 23)

	//fmt.Println(person.Greet())
	//person.HappyBirthday()
	//fmt.Println(person.Greet())

	var dog any = Dog{}
	checkType(dog)
}

type Dog struct{}
type Cat struct{}

func checkType(obj interface{}) {
	switch v := obj.(type) {
	case Dog:
		fmt.Println("obj adalah Dog:", v)
	case Cat:
		fmt.Println("obj adalah Cat:", v)
	default:
		fmt.Println("obj memiliki tipe lain")
	}
}
