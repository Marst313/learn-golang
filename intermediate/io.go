package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)

	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}

	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing from writer.", err)
	}
}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing from closer.", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // slack
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // heap
	_, err := buf.ReadFrom(mr)

	if err != nil {
		log.Fatalln("Error reading from reader.", err)
	}
	fmt.Println(buf.String())
}

func pipeExample() {
	pr, pw := io.Pipe()

	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func writeToFile(filepath string, data string) {
	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file.", err)
	}

	defer closeResource(file)

	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error opening/creating file.", err)
	}

	// Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error opening/creating file.", err)
	// }

}

func input_output() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello Reader !"))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer

	writeToWriter(&writer, "Hello Writer!")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	filepath := "io.txt"
	writeToFile(filepath, "Hello File!")

	resource := &MyResource{name: "Test resource"}
	closeResource(resource)
}

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing resource:", m.name)
	return nil
}

// ! WHEN TO USE IO PACKAGE
//  1. Purpose 				  = Basic interface and function for I/O operations
//  2. Core Tnterfaces/ Types = Core Interface: io.Reader, io.Writer, io.Closer, io.Seeker
//  3. Common Use Cases       = Direct I/o Operations, Simple Data Transfer, Pipes
//  4. Unique Use Cases       = Abstract I/O, Simple Data manipulation

// ! WHEN TO USE BUFIO PACKAGE
//  1. Purpose 				  = Buffered I/O opertaions to imporve efficiency bu reducing the number of system calls
//  2. Core Tnterfaces/ Types = Core Types: bufio.Reader, bufio.Writer, bufio.Scanner
//  3. Common Use Cases       = Buffered Reading/Writing, Line-by-Line Reading, Chunked Reading
//  4. Unique Use Cases       = Efficient I/O Operations, Complex Tokenization
