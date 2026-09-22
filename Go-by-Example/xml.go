// go offers built-in support for xml and xml-like formats with the encoding/xml package
package main

import (
	"encoding/xml"
	"fmt"
)

// plant will be mapped to xml
// similarly to the json examples, field tags contain directives for the encoder and decoder
// here we use some special features of the xml package: the XMLName field name dictates the name of the xml element representing this struct; id,attr means that the Id field is an xml attribure rather than a nested element
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id int `xml:"id,attr"`
	Name string `xml:"name"`
	Origin []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v", p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// emit xml representing our plant; using MarshalIndent to produce a more human-readable output
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// to add a generic xml header to the output, append it explicitlyf
	fmt.Println(xml.Header + string(out))

	// use Unmarshal to parse a stream of bytes with xml into a data structure
	// if the xml is malformed or cannot be mapped onto Plant, a descriptive error will be returned
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// the parent>child>plant field tag tells the encoder to nest all plants under <parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
