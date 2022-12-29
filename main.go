package main

import "fmt"

func main() {
	// This is Name Sugar variable syntax
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Println("Hello World")

	// using format Print in my output
	fmt.Printf("Welcome To %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Ticket Here To Attend: ")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name

	fmt.Println("Enter your first name :")
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastname")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n ", remainingTickets, conferenceName)
}

type Data struct {
	User   string
	Email  string
	Addres string
}
