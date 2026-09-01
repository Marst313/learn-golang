package main

import (
	"fmt"
)

func TypeConversion() {
	var a int = 32
	b := int32(a)
	c := float64(b)
	// d := bool("correct") //! Cant do this

	e := 3.14
	f := int(e)

	fmt.Println(f, c)

	//  HOW TO CONVER  <>Type(value)</>
	g := "Hello 😀 🇾🇹 übung"
	var h []byte
	h = []byte(g)

	fmt.Println("-------------------")
	fmt.Println(h)

	i := []byte{72, 101, 108, 108, 111}
	j := string(i)
	fmt.Println("-------------------")
	fmt.Println(j)

}
