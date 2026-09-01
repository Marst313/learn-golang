package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person2 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func struct_tags() {

	person := Person2{FirstName: "Karma", LastName: "Dharma", Age: 23}

	jsonData, err := json.Marshal(person)

	if err != nil {
		log.Fatalln("Error marshalling struct:", err)
	}

	fmt.Println(string(jsonData))

}
