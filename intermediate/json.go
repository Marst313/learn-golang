package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName    string  `json:"name"`
	Age          int     `json:"age,omitempty"`
	EmailAddress string  `json:"email,omitempty"`
	Address      Address `json:"address"`
}
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func json_tags() {

	person := Person{FirstName: "John"}

	//! Marshalling -> Convert GO data to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData))

	person1 := Person{FirstName: "Jane", Age: 30, EmailAddress: "jane@fakemail.com", Address: Address{
		City:  "Yogyakarta",
		State: "Indonesia",
	}}

	jsonData2, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println(string(jsonData2))

	jsonData1 := `{"full_name":"Jenny Doe", "emp_id":"0009", "age":30,"address": {"city":"Bantul","state":"Yogyakarta"}}`

	var employeeFromJSON Employee
	err = json.Unmarshal([]byte(jsonData1), &employeeFromJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println(employeeFromJSON)
	fmt.Println("Jenny's Age increased by 5 years", employeeFromJSON.Age+5)
	fmt.Println("Jenny's City", employeeFromJSON.Address.City)

	listOfCityState := []Address{
		{City: "New York", State: "NY"},
		{City: "San Jose", State: "CA"},
		{City: "Las Vegas", State: "NV"},
		{City: "Modesto", State: "CA"},
		{City: "Clearwater", State: "FL"},
	}

	fmt.Println(listOfCityState)
	jsonList, err := json.Marshal(listOfCityState)
	if err != nil {
		log.Fatalln("Error Marshalling to json")
	}

	fmt.Println("---------------- JSON LIST ------------------")
	fmt.Println("JSON List:", string(jsonList))

	fmt.Println("------------ UNKNOWN JSON STRUCTURE --------------------")
	// HANDLING Unknown JSON structure
	jsonData3 := `{"name":"John", "age":30 ,"address":{"city":"New York", "state":"NY"}}`
	var data map[string]any

	err = json.Unmarshal([]byte(jsonData3), &data)
	if err != nil {
		log.Fatalln("Error Unmarshalling JSON:", err)
	}

	fmt.Println("Decoded/Unmarshalled JSON:", data)
	fmt.Println("Decoded/Unmarshalled JSON:", data["address"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["name"])
	fmt.Println("Decoded/Unmarshalled JSON:", data["age"])

}

type Employee struct {
	FullName string  `json:"full_name"`
	EmpID    string  `json:"emp_id"`
	Age      int     `json:"age"`
	Address  Address `json:"address"`
}
