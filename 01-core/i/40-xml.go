package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	address Address  `xml:"address`
	// City    string   `xml:"city,omitempty"`
	Email string `xml:"-"`
	// Email   string   `xml:"email"`

}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	// person := Person{
	// 	Name:  "John",
	// 	Age:   30,
	// 	City:  "London",
	// 	Email: "example@email.com",
	// }
	person := Person{
		Name:  "John",
		Email: "example@email.com",
		address: Address{
			City:  "Oakland",
			State: "CA",
		},
	}
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error Marshaling data into XML", err)
	}
	fmt.Println("XML DATA:", string(xmlData))
	xmlData1, err := xml.MarshalIndent(person, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshaling data into XML", err)
	}
	fmt.Println("XML DATA1:", string(xmlData1))

	// Unmarshal

	// xmlRaw := `<person><name>Jane</name><age>25</age></person>`
	xmlRaw := `<person>
		<name>Jane</name>
		<age>25</age>
		<address>
			<city>
			""
			</city>
		</address>
	</person>`
	var personXML Person
	err = xml.Unmarshal([]byte(xmlRaw), &personXML)
	if err != nil {
		log.Fatalln("Error UnMarshaling XML", err)
	}
	fmt.Println("DATA:", personXML)
	fmt.Println("Local string:", personXML.XMLName.Local)
	fmt.Println("Namespace:", personXML.XMLName.Space)

	book := Book{
		ISBN:       "584-589-5245",
		Title:      "Go Learn",
		Author:     "Achyut",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attribute",
	}

	xmldataWithAttr, err := xml.MarshalIndent(book, "", "  ")
	if err != nil {
		log.Fatalln("Error Marshaling XML", err)
	}
	fmt.Println(string(xmldataWithAttr))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}

// <book isbn="asdasd" color="blue">
