package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
