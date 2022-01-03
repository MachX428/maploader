package util

import (
	"log"
	"os"
)

func InitLogging() {
	os.MkdirAll("maploader/log", os.ModePerm)
	err := RotateFile(7, "maploader/log/maploader", "log")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("maploader/log/maploader.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)
	log.Println("Logging initialized")
}