package main

import (
	"fmt"
)

func main () {
	var conferenceName = "Marks Conference"
	const confernceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Book yourself here, there are %v tickets and %v tickets left.\n", confernceTickets, remainingTickets)

	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask for their name
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email:")
	fmt.Scan(&email)
	fmt.Println("How many tickets would you like to purchase?")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for purchasing %v tickets! We will email you a confirmation to %v soon!\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v.\n", remainingTickets, conferenceName)

	fmt.Printf("These are all the bookings %v.\n", bookings)
}