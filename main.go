package main

import "fmt"

func main() {
	var userName string
	var age int
	var gender string
	var contactNumber string
	var userEmail string
	var fromStation string
	var toStation string
	var bookedTicketNumber int
	// var remainingTicket int
	// var totalTicket int
	// var trainName string
	// var seatType string
	// var trainDestination string
	// var trainDepartureTime string
	// var trainArrivalTime string
	// var ticketPrice float64
	// var seatNumber int
	// var seatBookedDate string
	// var seatBookedTime string

	// Replace the single Scan with individual prompts
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your gender: ")
	fmt.Scan(&gender)

	fmt.Print("Enter your contact number: ")
	fmt.Scan(&contactNumber)

	fmt.Print("Enter your email: ")
	fmt.Scan(&userEmail)

	fmt.Print("Enter your from station: ")
	fmt.Scan(&fromStation)

	fmt.Print("Enter your to station: ")
	fmt.Scan(&toStation)

	fmt.Print("Enter your booked ticket number: ")
	fmt.Scan(&bookedTicketNumber)

	fmt.Printf("\nBooking Information:\n")
	fmt.Printf("Name: %v\n", userName)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Gender: %v\n", gender)
	fmt.Printf("Contact number: %v\n", contactNumber)
	fmt.Printf("Email: %v\n", userEmail)
	fmt.Printf("From station: %v\n", fromStation)
	fmt.Printf("To station: %v\n", toStation)
	fmt.Printf("Booked ticket number: %v\n", bookedTicketNumber)
}
