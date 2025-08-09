package main

import "fmt"

func main() {
	var conferenceName = "Devanshi Godhaniya's"
	const conferenceTickets = 50
	var remainingTicket = 50

	fmt.Printf("Welcome to %v booking applicaiton\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are stil avilable.\n", conferenceTickets, remainingTicket)
	fmt.Println("Get your tickets here to attend")

}
