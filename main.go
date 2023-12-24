package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const tickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v Booking System \n", conferenceName)
	fmt.Println("Total tickets:", tickets)
	fmt.Println("Available tickets:", remainingTickets)
	fmt.Println("Get your ticket now!")

}
