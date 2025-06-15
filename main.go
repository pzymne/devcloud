package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

// var bookings []string
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	//var bookings [50]string

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			firstName := getFirstName()
			fmt.Printf("The first name of bookings is:%v\n", firstName)
			if remainingTickets == 0 {
				fmt.Println("our conference is booked out,come back next year")
				break
			}
		} else {
			fmt.Printf("Wo only have %v tickets remaining,so you can't book %v tickets\n", remainingTickets, userTickets)
			break
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still avable")
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var userTickets uint
	var email string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter your tickets num: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings is:%v\n", bookings)
	fmt.Printf("The whole slice:%v\n", bookings)
	fmt.Printf("The first value:%v\n", bookings[0])
	fmt.Printf("slice type:%T\n", bookings)
	fmt.Printf("slice length:%v\n", len(bookings))
	fmt.Printf("User %v booked %v tickets.You will receive conformation email at %v\n", firstName, userTickets, email)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("########################")
	fmt.Printf("Sending ticket:%v\n  to email address %v\n", ticket, email)
	fmt.Println("#######################")
}
