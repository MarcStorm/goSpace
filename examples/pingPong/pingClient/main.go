package main

import (
	"fmt"
	"goSpace/goSpace/topology"
	"goSpace/goSpace/tuplespace"
	"os"
	"strconv"
	"time"
)

func main() {
	// Get the port number that the user which to run the application on.
	pingPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	// Get the IP address and port number of the client running the pong
	// application.
	pongIP := os.Args[2]
	pongPort, err := strconv.Atoi(os.Args[3])
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	tsPtr := tuplespace.CreateTupleSpace(pingPort)
	fmt.Println(tsPtr)

	ownPtP, theirPtP := createPtP(pingPort, pongIP, pongPort)

	ping(ownPtP, theirPtP)
}

// Ping will initially act as client
func ping(ownPtP topology.PointToPoint, theirPtP topology.PointToPoint) {
	// Initialise the ping-pong by sending a "Ping" tuple to the pong
	// application's tuple space.
	for !tuplespace.Put(theirPtP, "Ping") {
		fmt.Println("Error in communication.")
	}
	fmt.Println("Ping Send")
	for {
		// Find a "Pong" tuple in own tuple space
		if tuplespace.Get(ownPtP, "Pong") {
			// A "Pong" tuple was found.
			fmt.Println("Pong recieved")

			// Send back a "Ping" tuple.
			if tuplespace.Put(theirPtP, "Ping") {
				fmt.Println("Ping Send")
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func createPtP(ownPort int, theirIP string, theirPort int) (topology.PointToPoint, topology.PointToPoint) {
	ownPtP := topology.CreatePointToPoint("Ping client", "localhost", ownPort)
	theirPtP := topology.CreatePointToPoint("Pong client", theirIP, theirPort)

	return ownPtP, theirPtP
}
