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
	var totalTicket int = 500
	var trainName string
	var departureTime string

	trainDepartureTime := []string{
		"06:15 AM",
		"11:15 AM",
		"03:45 PM",
		"08:00 PM",
		"10:30 PM",
	}
	trainOptions := []string{
		"Upobon Express",
		"Parabat Express",
		"Egarshindur",
		"Kishoreganj Express",
	}

	fmt.Println("\nAvailable Trains:")
	for i := 0; i < len(trainOptions); i++ {
		fmt.Printf("%d. %s - Departure: %s\n", i+1, trainOptions[i], trainDepartureTime[i])
	}
	var trainChoice int
	fmt.Print("Select train (enter number): ")
	fmt.Scan(&trainChoice)
	if trainChoice >= 1 && trainChoice <= len(trainOptions) {
		trainName = trainOptions[trainChoice-1]
		departureTime = trainDepartureTime[trainChoice-1]
		fmt.Printf("\nYou selected: %s\nDeparture Time: %s\n\n", trainName, departureTime)
	} else {
		fmt.Println("Invalid train selection. Please select a number between 1 and", len(trainOptions))
		return
	}

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
	fmt.Printf("Train Name: %v\n", trainName)
	fmt.Printf("Departure Time: %v\n", departureTime)
	fmt.Printf("Name: %v\n", userName)
	fmt.Printf("Age: %v\n", age)
	fmt.Printf("Gender: %v\n", gender)
	fmt.Printf("Contact number: %v\n", contactNumber)
	fmt.Printf("Email: %v\n", userEmail)
	fmt.Printf("From station: %v\n", fromStation)
	fmt.Printf("To station: %v\n", toStation)
	fmt.Printf("Booked ticket number: %v\n", bookedTicketNumber)
	remainingTickets := totalTicket - bookedTicketNumber
	if bookedTicketNumber <= totalTicket {
		fmt.Printf("Booking successful! Remaining tickets: %v\n", remainingTickets)
		totalTicket = remainingTickets
	} else {
		fmt.Printf("Sorry, only %v tickets are available\n", totalTicket)
	}
}
