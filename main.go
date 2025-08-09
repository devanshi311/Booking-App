package main

import "fmt"

func main() {
	conferenceName := "Devanshi Godhaniya's"
	const conferenceTickets int = 50
	var remainingTicket uint = 50

	fmt.Printf("Welcome to %v booking applicaiton\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil avilable.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for thier name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your first email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of the tickets:")
	fmt.Scan(&userTickets)

	remainingTicket = remainingTicket - userTickets

	fmt.Printf("Thank You %v %v for booking %v tickets.\nYou will receive a conformation email at %v \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTicket, conferenceName)
}
