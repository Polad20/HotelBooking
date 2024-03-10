package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userRooms uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidRoomNumber := userRooms > 0 && userRooms <= remainingRooms
	return isValidName, isValidEmail, isValidRoomNumber
}
