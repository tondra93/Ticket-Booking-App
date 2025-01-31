package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Booking struct {
	ID            int
	UserName      string
	Age           int
	Gender        string
	ContactNumber string
	Email         string
	FromStation   string
	ToStation     string
	TrainName     string
	DepartureTime string
	TicketNumber  int
	SeatNumbers   []string
}

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./railway.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create tables if they don't exist
	createTables := `
	CREATE TABLE IF NOT EXISTS bookings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name TEXT,
		age INTEGER,
		gender TEXT,
		contact_number TEXT,
		email TEXT,
		from_station TEXT,
		to_station TEXT,
		train_name TEXT,
		departure_time TEXT,
		ticket_number INTEGER,
		seat_number TEXT
	);

	CREATE TABLE IF NOT EXISTS trains (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		train_name TEXT UNIQUE,
		departure_time TEXT,
		total_seats INTEGER
	);`

	_, err = db.Exec(createTables)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func insertInitialTrains(db *sql.DB) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM trains").Scan(&count)
	if err != nil {
		log.Printf("Error checking trains: %v", err)
		return
	}

	// Only insert if no trains exist
	if count == 0 {
		trains := []struct {
			name          string
			departureTime string
		}{
			{"Upobon Express", "06:15 AM"},
			{"Parabat Express", "11:15 AM"},
			{"Egarshindur", "03:45 PM"},
			{"Kishoreganj Express", "08:00 PM"},
		}

		for _, train := range trains {
			_, err := db.Exec(`
				INSERT INTO trains (train_name, departure_time, total_seats)
				VALUES (?, ?, 500)
			`, train.name, train.departureTime)
			if err != nil {
				log.Printf("Error inserting train %s: %v", train.name, err)
			}
		}
	}
}

func displayMenu() int {
	fmt.Println("\n=== Railway Ticket Booking System ===")
	fmt.Println("1. Book a Ticket")
	fmt.Println("2. View Bookings")
	fmt.Println("3. Exit")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scan(&choice)
	return choice
}

func generateSeatNumbers(db *sql.DB, trainName string, numTickets int) ([]string, error) {
	// Get the last assigned seat number for this train
	var lastSeatNum int
	err := db.QueryRow(`
		SELECT COALESCE(MAX(CAST(REPLACE(REPLACE(seat_number, 'C', ''), 'W', '') AS INTEGER)), 0)
		FROM bookings 
		WHERE train_name = ?`, trainName).Scan(&lastSeatNum)
	if err != nil {
		return nil, err
	}

	var seatNumbers []string
	for i := 0; i < numTickets; i++ {
		lastSeatNum++
		// Alternate between window (W) and chair (C) seats
		seatType := "C"
		if lastSeatNum%2 == 0 {
			seatType = "W"
		}
		seatNumbers = append(seatNumbers, fmt.Sprintf("%s%d", seatType, lastSeatNum))
	}
	return seatNumbers, nil
}

func bookTicket(db *sql.DB) {
	var booking Booking

	// Get available trains - using a map to prevent duplicates
	rows, err := db.Query("SELECT DISTINCT train_name, departure_time FROM trains ORDER BY train_name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\nAvailable Trains:")
	var trains []struct {
		name string
		time string
	}

	// Use a map to track unique trains
	seenTrains := make(map[string]bool)

	for rows.Next() {
		var name, time string
		err := rows.Scan(&name, &time)
		if err != nil {
			log.Printf("Error scanning train row: %v", err)
			continue
		}

		// Only add if we haven't seen this train before
		if !seenTrains[name] {
			trains = append(trains, struct {
				name string
				time string
			}{name, time})
			seenTrains[name] = true
		}
	}

	// Display available trains
	for i, train := range trains {
		fmt.Printf("%d. %s - Departure: %s\n", i+1, train.name, train.time)
	}

	var trainChoice int
	fmt.Print("Select train (enter number): ")
	fmt.Scan(&trainChoice)

	if trainChoice < 1 || trainChoice > len(trains) {
		fmt.Println("Invalid train selection")
		return
	}

	booking.TrainName = trains[trainChoice-1].name
	booking.DepartureTime = trains[trainChoice-1].time

	fmt.Print("Enter your name: ")
	fmt.Scan(&booking.UserName)

	fmt.Print("Enter your age: ")
	fmt.Scan(&booking.Age)

	fmt.Print("Enter your gender: ")
	fmt.Scan(&booking.Gender)

	fmt.Print("Enter your contact number: ")
	fmt.Scan(&booking.ContactNumber)

	fmt.Print("Enter your email: ")
	fmt.Scan(&booking.Email)

	fmt.Print("Enter your from station: ")
	fmt.Scan(&booking.FromStation)

	fmt.Print("Enter your to station: ")
	fmt.Scan(&booking.ToStation)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&booking.TicketNumber)

	// Generate seat numbers
	seatNumbers, err := generateSeatNumbers(db, booking.TrainName, booking.TicketNumber)
	if err != nil {
		log.Printf("Error generating seat numbers: %v", err)
		fmt.Println("Failed to generate seat numbers")
		return
	}
	booking.SeatNumbers = seatNumbers

	// Modify the INSERT query to include seat numbers
	for _, seatNum := range booking.SeatNumbers {
		_, err = db.Exec(`
			INSERT INTO bookings (
				user_name, age, gender, contact_number, email,
				from_station, to_station, train_name, departure_time, 
				ticket_number, seat_number
			) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			booking.UserName, booking.Age, booking.Gender, booking.ContactNumber,
			booking.Email, booking.FromStation, booking.ToStation, booking.TrainName,
			booking.DepartureTime, 1, seatNum)

		if err != nil {
			log.Printf("Error booking ticket: %v", err)
			fmt.Println("Failed to book ticket")
			return
		}
	}

	fmt.Println("\nBooking successful!")
	displayBookingInfo(booking)
}

func viewBookings(db *sql.DB) {
	rows, err := db.Query(`
		SELECT user_name, train_name, departure_time, seat_number 
		FROM bookings
		ORDER BY train_name, seat_number
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n=== All Bookings ===")
	fmt.Printf("%-20s %-20s %-15s %-10s\n", "Name", "Train", "Departure", "Seat")
	fmt.Println(strings.Repeat("-", 65))

	for rows.Next() {
		var name, train, departure, seat string
		rows.Scan(&name, &train, &departure, &seat)
		fmt.Printf("%-20s %-20s %-15s %-10s\n", name, train, departure, seat)
	}
}

func displayBookingInfo(booking Booking) {
	fmt.Printf("\nBooking Information:\n")
	fmt.Printf("Train Name: %v\n", booking.TrainName)
	fmt.Printf("Departure Time: %v\n", booking.DepartureTime)
	fmt.Printf("Name: %v\n", booking.UserName)
	fmt.Printf("Age: %v\n", booking.Age)
	fmt.Printf("Gender: %v\n", booking.Gender)
	fmt.Printf("Contact number: %v\n", booking.ContactNumber)
	fmt.Printf("Email: %v\n", booking.Email)
	fmt.Printf("From station: %v\n", booking.FromStation)
	fmt.Printf("To station: %v\n", booking.ToStation)
	fmt.Printf("Number of tickets: %v\n", booking.TicketNumber)
	fmt.Printf("Seat Numbers: %v\n", strings.Join(booking.SeatNumbers, ", "))
}

func main() {
	db := initDB()
	defer db.Close()

	insertInitialTrains(db)

	for {
		choice := displayMenu()

		switch choice {
		case 1:
			bookTicket(db)
		case 2:
			viewBookings(db)
		case 3:
			fmt.Println("Thank you for using our service!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
