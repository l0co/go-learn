package main

import "fmt"

// structs are not inheritable, but can be composited

type structa struct {
	a string
}

type structb struct {
	b string
}

// structc is composed of structa and structb what means it gets the access to all their fields directly
type structc struct {
	structa
	structb
}

func main() {

	{
		a := structa{"a"}
		b := structb{"b"}
		c := structc{a, b}
		fmt.Printf("%+v\n", c) // {structa:{a:a} structb:{b:b}}

		// we can access composition struct fields directly from composited struct
		fmt.Println(c.a) // a

		// if the composition struct changes, it's not reflected in composited struct,
		// so for composited struct a new composition struct instance is created
		a.a = "A"
		fmt.Printf("%+v\n", a) // {a:A}
		fmt.Printf("%+v\n", c) // {structa:{a:a} structb:{b:b}}

		// the way to change values
		c.a = "A"
		c.structb.b = "B"
		fmt.Printf("%+v\n", c) // {structa:{a:A} structb:{b:B}}
	}

	// so, the real way we can have separated composite struct instances which can be changed independently is the following:
	{
		c := structc{structa{"a"}, structb{"b"}}
		a := &c.structa
		b := &c.structb

		a.a = "A"
		b.b = "B"
		fmt.Printf("%+v\n", c) // {structa:{a:A} structb:{b:B}}
	}

}
