package main

import (
	"encoding/xml"
	"fmt"
)

// same as for json, we use structs with tags
type Plant struct {
	fmt.Stringer          // this interface is used by fmt.Print to get string version of the object
	XMLName      xml.Name `xml:"plant"`
	Id           int      `xml:"id,attr"`
	Name         string   `xml:"name"`
	Origin       []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// simple marshalling
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	// <plant id="27">
	//   <name>Coffee</name>
	//   <origin>Ethiopia</origin>
	//   <origin>Brazil</origin>
	// </plant>

	// with xml header
	fmt.Println(xml.Header + string(out))
	// <?xml version="1.0" encoding="UTF-8"?>
	//  <plant id="27">
	//   <name>Coffee</name>
	//   <origin>Ethiopia</origin>
	//   <origin>Brazil</origin>
	//  </plant>

	// unmarshalling
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p) // Plant id=27, name=Coffee, origin=[Ethiopia Brazil]

	// nesting another level of xml
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
	// <nesting>
	//   <parent>
	//     <child>
	//       <plant id="27">
	//         <name>Coffee</name>
	//         <origin>Ethiopia</origin>
	//         <origin>Brazil</origin>
	//       </plant>
	//       <plant id="81">
	//         <name>Tomato</name>
	//         <origin>Mexico</origin>
	//         <origin>California</origin>
	//       </plant>
	//     </child>
	//   </parent>
	// </nesting>
}
