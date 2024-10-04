package utils

import (
	"log"
	"os"
)

var (
	logger = log.New(os.Stdout, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func Info(message string) {
	logger.Println("INFO: " + message)
}

func Error(message string) {
	logger.Println("ERROR: " + message)
}
