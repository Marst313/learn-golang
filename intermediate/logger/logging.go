package main

import (
	"fmt"
	"log"
	"os"
)

func logger() {

	log.Println("This is a log message.")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message.")

	//! LOG FLAGS
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("This is a log message with a date, time.")

	fmt.Println("-----------------------------------------")

	infoLogger.Println("This is an info message.")
	warnLogger.Println("This is a warning message.")
	errorLogger.Println("This is an error message")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file :%v", err)
	}
	defer file.Close()

	//! WRITE IT IN FILE
	infoLogger1 := log.New(file, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger1 := log.New(file, "WARN:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger1 := log.New(file, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

	debugLogger.Println("This is a debug message.")

	warnLogger1.Println("This is a warning message.")
	infoLogger1.Println("This is a info message.")
	errorLogger1.Println("This is an error.")

}

// ! LOG ONLY IN TERMINAL
var (
	infoLogger  = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN:", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
)
