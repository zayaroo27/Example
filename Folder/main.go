package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
	func main() {
		confrenceName := "GO confrence"
		remainingTickets := 30
		const conferenceTickets uint = 30
		var bookings []string
		fmt.Printf("Welcome to %v \nwrite your confrence name\n", confrenceName)
		fmt.Printf("You have %v tickets remaining out of %v\n", remainingTickets, conferenceTickets)

		for {
			var firstName string
			var lastName string
			var userTickets int
			var userEmail string
			fmt.Println("Enter your first name")
			fmt.Scan(&firstName)
			fmt.Println("Enter your last name")
			fmt.Scan(&lastName)
			fmt.Println("Enter your Email")
			fmt.Scan(&userEmail)
			fmt.Println("Enter the amount of tickets")
			fmt.Scan(&userTickets)
			if userTickets > remainingTickets {
				fmt.Printf("Sorry, We only have %v tickets remaining, You can't book %v tickets", remainingTickets, userTickets)
				break
			}
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for buying %vtickets, We'll send a confirmation letter to %v\n", firstName, lastName, userTickets, userEmail)
			fmt.Printf("We have %v remaining tickets\n", remainingTickets)
			fmt.Printf("These are our bookings %v\n", bookings)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)

				firstNames = append(firstNames, names[0])

			}
			fmt.Printf("These are the first names of bookings %v\n", firstNames)
			noremainingTickets := remainingTickets == 0

			if noremainingTickets {
				fmt.Println(" The tickets are all sold out!")
				break
			}

		}
	}
*/

func userInput(input *int) {
	*input = 32

}
func addNum(x int, y int) (int, int, error) {
	if y == 0 {
		return 0, 0, errors.New("error: eivision by zero is not allowed")

	}
	result := x / y
	reamainder := x % y
	return result, reamainder, nil

}

func timeLoop(slice []int, n int) time.Duration {
	var Time0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)

	}
	return time.Since(Time0)
}

type myEngine struct {
	mpg     uint8
	gallons uint8
}
type myElectricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (x myEngine) milesleft() uint8 {
	return x.mpg * x.gallons
}
func (x myElectricEngine) milesleft() uint8 {
	return x.mpkwh * x.kwh
}

type FirstStruct struct {
	Name string
	Age  uint8
	SecondStruct
	int
}
type engine interface {
	milesleft() uint8
}
type SecondStruct struct {
	Job string
}
type RudeWhite struct {
	Name string
	age  uint8
}
type RudeBlack struct {
	BlackName string
	president bool
}

func (r RudeWhite) sayCurseWord() bool {
	if r.age >= 18 {
		return true

	} else {
		return false
	}

}
func (r RudeBlack) sayCurseWord() bool {
	if r.president {
		fmt.Println(r.BlackName)
		return true

	} else {
		fmt.Println(r.BlackName)
		return false
	}
}

func (r *RudeBlack) BePresident(name string) {
	r.BlackName = name
}

type Race interface {
	sayCurseWord() bool
}

func Racism(r Race, word string) {
	if r.sayCurseWord() {
		fmt.Println(word)

	} else {
		fmt.Printf("Shut your ass")
	}
}

func canMakeIt(e engine, miles uint8) {
	if e.milesleft() >= miles {
		fmt.Println("You can make it")
	} else {
		fmt.Println("You are out of fuel")
	}
}

func main() {
	// numarator := 24
	// denomator := 0
	// result, remainder, err := addNum(numarator, denomator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Println("The number is divisible")
	// } else {
	// 	fmt.Printf("The result of the division is %v and the reaminder is %v\n", result, remainder)
	// }
	// n := 1000000
	// var mySlice = []int{}
	// var mySlice2 = make([]int, 0, n)
	// fmt.Printf("Duration without preset %v \n", timeLoop(mySlice, n))
	// fmt.Printf("Duration with preset %v", timeLoop(mySlice2, n))
	var Array = []string{"A", "B", "c", "d", "e"}
	var StrBuilder strings.Builder
	var myDetail FirstStruct = FirstStruct{
		Name:         "Zayar",
		Age:          19,
		SecondStruct: SecondStruct{Job: "IT manager"},
	}
	var firstEngine engine = myEngine{34, 3}
	var secondEngine myElectricEngine
	var White RudeWhite = RudeWhite{"Eminum", 18}
	var Black RudeBlack = RudeBlack{"Drake", false}
	Black.BePresident("Not Drake")
	Racism(White, "Nigga")
	Racism(Black, "Pig")
	secondEngine.kwh = 3
	secondEngine.mpkwh = 5
	for i := range Array {
		StrBuilder.WriteString(Array[i])

	}
	BuiltStr := StrBuilder.String()
	fmt.Println(BuiltStr)

	myDetail.Age = 23
	fmt.Println(myDetail.Name, myDetail.Age, myDetail.Job, myDetail.int)

	canMakeIt(firstEngine, 100)
	canMakeIt(secondEngine, 32)

}
