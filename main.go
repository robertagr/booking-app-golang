package main

import "fmt"


func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50 
	var remainingTickets uint = 50


	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var userName string
	var userTickets int
	// ask user for their name
	
	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked this %v tickets\n", userName, userTickets)
}
