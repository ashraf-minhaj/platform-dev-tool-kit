/*
Building a developer platform

author: ashraf-minhaj
mail: ashraf_minhaj@yahoo.com
*/

package main

import (
	"fmt"
	"os"
	"strings"
)

func ticket_processor() {

}

func docString() {
	fmt.Println("Usage: platform start ticket EPD-123")
}

func isValidTicket(ticket string) bool {
	// ticketStartsWith := []string{"EP", "EPD"}

	parts := strings.Split(ticket, "-")
	if len(parts) == 2 {
		if parts[0] == "EP" || parts[0] == "EPD" {
			// if strconv.Atoi(parts[1])
			return true
		}
	}

	fmt.Println(parts)

	return false
}

func main() {
	isWrongCommand := false

	fmt.Println("Hello, Platform user!")
	fmt.Println("If you face any difficulty feel free to reach out to the platform team")

	arguments := os.Args
	fmt.Println(arguments)

	// check number of arguments
	if len(arguments) < 4 {
		isWrongCommand = true
	} else if arguments[2] == "ticket" {
		ticket := os.Args[3]
		if isValidTicket(ticket) {
			fmt.Println(ticket)
		} else {
			isWrongCommand = true
		}
	} else {
		isWrongCommand = true
	}

	// fmt.Println(isWrongCommand)
	if isWrongCommand {
		docString()
		return
	}
}
