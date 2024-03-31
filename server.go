package main

import (
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {

	// init logger
	initLogger()

	// read configuration
	/*
		As configuration info we will use only number on which port the server will start
		It will be sent to server as first argument in cli
	*/
	port, err := ReadPortNumber()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("Server will run on localhost:%s\n", port)

	// run server

}

// Create or Open file logs for storing all logs in it
func initLogger() {
	file, err := os.OpenFile("logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Fatal("Failed to log to file, using default stderr")
	}
}

// Reads port number from programm arguments
func ReadPortNumber() (s string, e error) {

	// Check if we have an arguments -> Args[0] is the program name
	if len(os.Args) == 1 {
		e = errors.New("Not enough arguments to start program!\nPlease specify on which port server should be hosted\n")
		return "", e
	}

	// If the 1st argument exists
	s = os.Args[1]

	// Try to convert argument value to Int
	port, e := strconv.Atoi(s)

	// If fail - incorrect type
	if e != nil {
		e = errors.New("Invalid type of argument!\nPlease write number on which port server should be hosted\n")
		return "", e
	}

	// Check if number is in range [1, 65535]
	if port <= 0 || port > 65535 {
		e = errors.New("Invalid value of port number!\nPlease set port number from range [1, 65535]\n")
		return "", e
	}

	// Now all validations are passed, so port number can be used in porogram
	return s, nil
}
