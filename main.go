package main

import (
	"fmt"
	"strings"
)



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
		
		// ask for user input 
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)
		
		isValidName := len(firstName) >= 2 && len(lastName) >=2
        isValidEmail := strings.Contains(email, "@")  //if email contains @
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

        // isValidCity := city == "Singapore" || city == "London"


		if isValidEmail && isValidName && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			//SLICE
			bookings = append(bookings, firstName + " " + lastName)
		
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
		
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking) 
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
		
			if remainingTickets == 0 {
				// then, end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("first name of last name you entered is too short")
			} 
			if !isValidName {
				fmt.Println("email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}
	
	city := "London"

	switch city {
	  case "New York":
		// execute code for booking NY conference tickets
	  case "Singapore", "Hong Kong":
		// execute code for booking Singapore and HK conference tickets
	  case "London", "Berlin":
       // some code here for London and Berlin
	  case "Mexico":
      // some code here

	default:
		fmt.Print("No valid city selected")

	}
}

func greetUsers() {
	
}