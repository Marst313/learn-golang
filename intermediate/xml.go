package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type PersonXML struct {
	XMLName xml.Name   `xml:"person"`
	Name    string     `xml:"name"`
	Age     int        `xml:"age"`
	Email   string     `xml:"-"`
	Address AddressXML `xml:"address"`
	// Email   string   `xml:"email"`
}

type AddressXML struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func XML() {

	person := PersonXML{Name: "Karma", Age: 20, Email: "Karmadharma@gmail.com", Address: AddressXML{City: "YG", State: "ID"}}

	//! WITHOUT INDENTATION
	xmlData, err := xml.Marshal(person)
	if err != nil {
		log.Fatalln("Error marshalling data into xml: ", err)
	}
	fmt.Println(string(xmlData))

	//! WITH INDENTATION
	xmlData1, err := xml.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling data into xml: ", err)
	}
	fmt.Println(string(xmlData1))

	//! UNMARSHALLING

	rawData := `<person><name>Dharma</name><age>25</age><address><city>WNG</city><state>ID</state></address></person>`

	var personXML PersonXML

	err = xml.Unmarshal([]byte(rawData), &personXML)
	if err != nil {
		log.Fatalln("Error Unmarshalling data into xml: ", err)
	}

	fmt.Println("--------------------------------")
	fmt.Println(personXML)
	fmt.Println(personXML.XMLName.Local, "Local string")
	fmt.Println(personXML.XMLName.Space, "Namespace")

	book := Book{
		ISBN:       "585-2423-2323-222",
		Title:      "My Name is methos",
		Author:     "Methos",
		Pseudo:     "Pseudo",
		PseudoAttr: "Pseudo Attritbute",
	}

	xmlDataAttr, err := xml.MarshalIndent(book, "", " ")
	if err != nil {
		log.Fatalln("Error marshalling to xml", err)
	}
	fmt.Println("--------------------------------")
	fmt.Println(string(xmlDataAttr))

}

type Book struct {
	XMLName    xml.Name `xml:"book"`
	ISBN       string   `xml:"isbn,attr"`
	Title      string   `xml:"title,attr"`
	Author     string   `xml:"author,attr"`
	Pseudo     string   `xml:"pseudo"`
	PseudoAttr string   `xml:"pseudoattr,attr"`
}
