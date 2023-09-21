package main

import "fmt"


func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50 
	var remainingTickets uint = 50
	//ARRAY
	// var bookings [50]string

	//SLICE
	var bookings []string
	//alternative syntax for slice:
	// var bookings []string{}
	// bookings := []string{}


	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	
	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		
		remainingTickets = remainingTickets - userTickets
	
		//ARRAY
		// bookings[0] = firstName + " " + lastName  
	
		//SLICE
		bookings = append(bookings, firstName + " " + lastName)
	
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	
		fmt.Printf("These are all the bookings: %v\n", bookings)

	}
	

}
