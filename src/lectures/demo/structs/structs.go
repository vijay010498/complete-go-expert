package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	vijay := Passenger{"vijay", 1, false}
	fmt.Println(vijay)

	var (
		pablo = Passenger{Name: "Pablo", TicketNumber: 2}
	)

	var neha Passenger
	neha.Name = "Neha"
	neha.TicketNumber = 4
	fmt.Println(neha)

	vijay.Boarded = true
	pablo.Boarded = true

	if vijay.Boarded {
		fmt.Println(vijay.Name, "has Boarded the bus")
	}
	bus := Bus{vijay}
	fmt.Println("Bus is here ", bus, "in front seat", bus.FrontSeat.Name)
}
