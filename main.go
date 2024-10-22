package main

import (
	"fmt"
	"strings"
)

func main() {
	confrenceName := "Go Confrence"
	const confrenceTickets int = 50

	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application.\n", confrenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confrenceTickets, remainingTickets)
	fmt.Printf("Get your %v tickets here to attend.\n", confrenceName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// asks user for their name
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)

		// asks user for email
		fmt.Print("Enter your email: ")
		fmt.Scan(&email)

		// ask user for how many tickets they want
		fmt.Print("How many tickets do you want to buy? ")
		fmt.Scan(&userTickets)

		len(firstName) >= 2

		if userTickets <= int(remainingTickets) {
			remainingTickets -= uint(userTickets)
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("We have %v tikckets remaining for %v\n", remainingTickets, confrenceName)

			firstNames := []string{}
			for _, booking := range bookings { // "_" a balnk identifier is used to ignore a variable you don't want to use
				var names = strings.Fields(booking)
				// strings.Field() splits the string with white space as seperator
				// var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}
			// the range function iterates over differnt elemsnts fo diferent dtata structures
			// including arrays, slices, etc.
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			var notTicketsRemaining bool = remainingTickets == 0 // or

			if notTicketsRemaining {
				fmt.Println("Tickets are sold out! Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remianing, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue // using a break statement ends the program abruptly without giving users a chance to make correct input

		}

	}

}
