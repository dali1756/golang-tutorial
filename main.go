package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("hello world.")
	fmt.Println(strings.Repeat("-", 50))

	// variable and constants
	var userName string = "Jeter"
	fmt.Println(userName)
	fmt.Println(strings.Repeat("-", 50))

	const number int = 123
	fmt.Println(number)
	fmt.Println(strings.Repeat("-", 50))

	// formatted
	var userName2 = "Ray"
	const years = 10
	fmt.Println("name:", userName2, "years old:", years)
	fmt.Printf("hello my name is %v, I'm %v years old.\n", userName2, years)
	fmt.Println(strings.Repeat("-", 50))

	// data type
	var userName3 string
	var number2 int
	userName3 = "hello world."
	number2 = 123
	fmt.Printf("name: %v number: %v\n", userName3, number2)
	fmt.Println(strings.Repeat("-", 50))

	// float
	var price float64 = 199.5
	var rate float32 = 0.05
	discount := 0.8
	fmt.Printf("%T %T %T\n", price, rate, discount)
	fmt.Println(strings.Repeat("-", 50))

	// int and float 需要明確轉換
	tickets := 3
	total := price * float64(tickets)
	fmt.Println("total:", total)
	fmt.Println(strings.Repeat("-", 50))

	// 整數除法和浮點數除法
	fmt.Println(7/2, 7.0/2, float64(7)/2)
	fmt.Println(strings.Repeat("-", 50))

	// 格式化
	pi := math.Pi
	fmt.Printf("%v | %.2f | %8.3f | %e\n", pi, pi, pi, pi)
	fmt.Println(strings.Repeat("-", 50))

	// precision
	a, b := 0.1, 0.2
	fmt.Println(a+b, a+b == 0.3)
	fmt.Println(math.Abs(a+b-0.3) < 1e-9)
	fmt.Println(strings.Repeat("-", 50))

	// float32 vs float64
	var f32 float32 = 1.0 / 3
	var f64 float64 = 1.0 / 3
	fmt.Println(f32, f64)
	fmt.Println(strings.Repeat("-", 50))

	// 把浮點數轉成整數會截斷
	x, y := 3.99, -3.99
	fmt.Println(int(x), int(y))
	fmt.Println(math.Round(3.5), math.Floor(3.99), math.Ceil(3.01))
	fmt.Println(strings.Repeat("-", 50))

	// 特殊東西
	zero := 0.0
	fmt.Println(1/zero, -1/zero, zero/zero, math.IsNaN(zero/zero))
	fmt.Println(strings.Repeat("-", 50))

	test := 123.4
	fmt.Printf("%T\n", test)
	fmt.Println(strings.Repeat("-", 50))

	// ask user for their name
	var userName4 string
	fmt.Println("請輸入你的名字：")
	fmt.Scan(&userName4)
	fmt.Println("你好，", userName4)
	fmt.Println(strings.Repeat("-", 50))

	// array
	// var bookings [50]string
	// slice
	// var bookings []string
	// array 做法
	// bookings[0] = firstName + " " + lastName
	// slice 做法
	// bookings = append(bookings, firstName + " " + lastName)

	// fmt.Printf("the whole array：%v\n", bookings)
	// fmt.Printf("the first value：%v\n", bookings[0])
	// fmt.Printf("array type：%T\n", bookings)
	// fmt.Printf("array length：%v\n", len(bookings))

	// fmt.Printf("the whole slice：%v\n", bookings)
	// fmt.Printf("the first value：%v\n", bookings[0])
	// fmt.Printf("slice type：%T\n", bookings)
	// fmt.Printf("slice length：%v\n", len(bookings))

	// array and slice
	// var bookings = [50]string{"Jeter", "Jay"}

	// 以下兩種方式都可以
	// var bookings = [50]string{}
	var bookings1 [50]string
	bookings1[0] = "Jeter"
	bookings1[1] = "Jay"
	fmt.Printf("%v\n", bookings1)
	fmt.Println(strings.Repeat("-", 50))

	// loops
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%v X %v = %2v  ", i, j, i*j)
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("-", 50))

	// if-else and bool
	score := 60
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	fmt.Println(strings.Repeat("-", 50))

	// switch
	city := "Taiwan"
	switch city {
	case "New York":
		fmt.Println("you choices New York.")
	case "Singapore", "Hong kong":
		fmt.Println("you choices singapore or hong kong.")
	case "London", "Berlin":
		fmt.Println("you choices london or berlin.")
	default:
		fmt.Println("no valid city selected.")
	}

	// ex. book ticket logic
	conferenceName := "Jeter"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	for {
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

		// user input validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("thank you %v %v for booking %v tickets. you will reveive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// For-Each
			// firstNames := []string{}
			// 因 := 是空的沒有資料所以按照慣例使用 var 會比較適合
			var firstNames []string
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("the first name of bookings are：%v\n", firstNames)
			fmt.Println(strings.Repeat("-", 50))

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
