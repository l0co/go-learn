package main

import (
	"fmt"
	"reflect"
)

// go has no classes but structs instead (but no inheritance)

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	// struct initialization
	p := person{name: name}
	p.age = 42
	// normally pointers to struct are formatted as &{...} with struct content,
	// this does the job of displaying real addresses
	fmt.Printf("newPersonByPointer: %p\n", &p)
	// that's odd because we return a pointer to local function variable, but it apparently survives
	// so in go local function variables aren't destroyed when function exits, but probably when GC sees no references to variable
	return &p
}

func newPersonVal(name string) person {
	p := person{name: name}
	p.age = 42
	fmt.Printf("newPersonByVariable: %p\n", &p)
	// we can also return a full variable here
	return p
}

func main() {

	fmt.Println(person{"Bob", 20})              // initialization with consecutive field values: {Bob 20}
	fmt.Println(person{name: "Alice", age: 30}) // initialization with explicite field values: {Alice 30}
	fmt.Println(person{name: "Fred"})           // not all fields needs to be initialized, they will be zero-valued: {Fred 0}
	fmt.Println(&person{name: "Ann", age: 40})  // a struct pointer can be created immediately: &{Ann 40}

	// you don't have to create a new person when variable is declared to fill it with zero-valued value (
	// struct instance with all fields zero-valued)
	var myPerson person
	fmt.Println("myPerson", myPerson) // myPerson { 0}

	// however the uninitialized pointer is empty (nil) and no allocation is done
	var myPersonP *person
	fmt.Println("myPersonP", myPersonP) // myPerson <nil>

	{
		p := newPerson("Jon")
		fmt.Println(p)
		fmt.Printf("local newPersonByPointer: %p\n", p)
		// newPersonByPointer: 0xc00000c100
		// &{Jon 42}
		// local newPersonByPointer: 0xc00000c100
		// recap: we created the variable in local function and returned the pointer to it, here it's the same variable as locally in function
	}
	{
		p := newPersonVal("Jon")
		fmt.Println(p)
		fmt.Printf("local newPersonByVariable: %p\n", &p)
		// newPersonByVariable: 0xc00009a120
		// {Jon 42}
		// local newPersonByVariable: 0xc00009a100
		// recap: we created the variable in local function and returned it, what allocated a new variable which is accessible here now
	}

	// accessing the struct field
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// same possible for struct pointers
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 51
	fmt.Println(sp.age)

	// struct can be defined anonymously and immediately instantiated
	x := struct {
		a int
	}{1}
	fmt.Println(x) // {1}

	// struct field can have tags
	{
		type taggedStruct struct {
			// F with capital letter is exported, // otherwise idea complains about unexported field with a tag
			F string `someTag:"value" someOtherTag:"otherValue"` // this field has two tags
		}

		t := taggedStruct{F: "f"}
		taggedStructType := reflect.TypeOf(t)
		f, _ := taggedStructType.FieldByName("F")
		fmt.Println(f.Name)                                                                       // F
		fmt.Println(f.Type)                                                                       // string
		fmt.Println(f.Tag)                                                                        // someTag:"value" someOtherTag:"otherValue"
		fmt.Println(f.Tag.Get("someTag"), f.Tag.Get("someOtherTag"), f.Tag.Get("nonExistingTag")) // value otherValue (+empty string)
	}

}
