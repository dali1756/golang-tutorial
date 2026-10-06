package main

import (
	"fmt"
	"golang-tutorial/helper"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Jeter"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// ex. book ticket
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
            go sendTicket(userTickets,firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("the first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("our conference is book out. come back next year.")
				break
			}
			// 防止使用者輸入票數大於總票數
		} else {
			if !isValidName {
				fmt.Println("請輸入正確的 first name or last name。")
			}
			if !isValidEmail {
				fmt.Println("您的 email 沒有包含 @ 字元。")
			}
			if !isValidTicketNumber {
				fmt.Println("您輸入的票數無效。")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("welcome to %v booking application.\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("get your tickets here to attend.")
}

func getFirstNames() []string {
	// For-Each
	// firstNames := []string{}
	// 因 := 是空的沒有資料所以按照慣例使用 var 會比較適合
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("請輸入您的名字：")
	fmt.Scan(&firstName)

	fmt.Println("請輸入您的姓氏：")
	fmt.Scan(&lastName)

	fmt.Println("請輸入您的 email：")
	fmt.Scan(&email)

	fmt.Println("請輸入您的票：")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	// create a map for a user
	userData := userData {
        firstName: firstName,
        lastName: lastName,
        email: email,
        numberOfTickets: userTickets,
    }
	bookings = append(bookings, userData)

	fmt.Printf("list of bookings is %v\n", bookings)
	fmt.Printf("thank you %v %v for booking %v tickets. you will reveive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
    time.Sleep(10 * time.Second)
    var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
    fmt.Println("====================")
    fmt.Printf("sending ticket：\n %v \nto email address %v\n", ticket, email)
    fmt.Println("====================")
}