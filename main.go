package main

import (
	"fmt"
	//"sync"
	"time"
)

const totalRooms int = 50

var applicationName = "Hotel booking"
var remainingRooms uint = 50 // number of remaining rooms can`t be negative
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	numberOfRooms uint
}

//var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userRooms := getUserInput()
		isValidName, isValidEmail, isValidRoomNumber := validateUserInput(firstName, lastName, email, userRooms)

		if isValidName && isValidEmail && isValidRoomNumber {

			bookRoom(userRooms, firstName, lastName, email)

			//wg.Add(1) if we don`t use For loope
			go sendInfo(userRooms, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingRooms == 0 {
				fmt.Println("We don`t have so many rooms. Try less or come back next month.")
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidRoomNumber {
				fmt.Println("number of rooms you entered is invalid")
			}
		}
	}
	//wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v application\n", applicationName)
	fmt.Printf("We have total of %v rooms and %v are still available.\n", totalRooms, remainingRooms)
	fmt.Println("Book a room here to check in")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userRooms uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of rooms: ")
	fmt.Scan(&userRooms)

	return firstName, lastName, email, userRooms
}

func bookRoom(userRooms uint, firstName string, lastName string, email string) {
	remainingRooms = remainingRooms - userRooms

	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		numberOfRooms: userRooms,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v rooms. You will receive a confirmation email at %v\n", firstName, lastName, userRooms, email)
	fmt.Printf("%v rooms remaining for booking\n", remainingRooms)
}

func sendInfo(userRooms uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var room = fmt.Sprintf("%v rooms for %v %v", userRooms, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending info:\n %v \nto email address %v\n", room, email)
	fmt.Println("#################")
	//wg.Done()
}
