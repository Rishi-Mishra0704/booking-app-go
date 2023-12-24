package main

import (
	"fmt"
	"sync"
	"time"
)

// Constants for the conference
const (
	conferenceTickets int = 50
	conferenceName        = "Go Conference"
)

// Global variables
var (
	remainingTickets uint           = 50
	bookings         []UserData     // Slice to store user bookings
	wg               sync.WaitGroup // WaitGroup to wait for goroutines to finish
)

// UserData struct represents user information
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// main function is the entry point of the program
func main() {
	// Greet users and display available tickets
	greetUsers()

	// Infinite loop (commented out) - User can book tickets until the conference is booked out
	// for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	// Check if user input is valid
	if isValidName && isValidEmail && isValidTicketNumber {
		// Book the tickets and send confirmation
		bookTicket(userTickets, firstName, lastName, email)

		// Start a goroutine to simulate sending a ticket (async operation)
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Display the first names of all bookings
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		// Check if all tickets are booked, end the program
		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		// Print error messages for invalid input
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	//}
	wg.Wait() // Wait for all goroutines to finish before exiting
}

// greetUsers function displays a welcome message and available tickets information
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// getFirstNames function returns a slice of first names from all bookings
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// getUserInput function takes input from the user and returns the values
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// bookTicket function books tickets for the user and updates global variables
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Create UserData struct and append to bookings slice
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)

	// Display booking information
	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// sendTicket function simulates sending a ticket after a delay (async operation)
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	// Display a message indicating that the ticket is being sent
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")

	// Notify WaitGroup that the goroutine has finished
	wg.Done()
}
