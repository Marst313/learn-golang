package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func logrus_logger() {

	log := logrus.New()

	//! SET LOG LEVEL
	log.SetLevel(logrus.InfoLevel)

	//! SET LOG FORMAT
	log.SetFormatter(&logrus.JSONFormatter{})

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Printf("Erorr opening file %v", err)
	}
	defer file.Close()

	// ! LOGGING EXAMPLES
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")
}
