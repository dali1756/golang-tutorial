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
	fmt.Println(a + b, a + b == 0.3)
	fmt.Println(math.Abs(a + b - 0.3) < 1e-9)
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

    // ex. book ticket logic
    conferenceName := "Jeter"
    const conferenctTickets int = 50
    var remainingTickets uint = 50
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

    remainingTickets = remainingTickets - userTickets

    fmt.Printf("thank you %v %v for booking %v tickets. you will reveive a confirmation email at %v\n", firstName, lastName, userTickets, email)
    fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
    fmt.Println(strings.Repeat("-", 50))
}
