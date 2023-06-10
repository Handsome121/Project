package main

import (
	"log"
	"time"
)

//简单定时任务
func TicketDemo() {
	ticket := time.NewTicker(1 * time.Second)
	defer ticket.Stop()

	for range ticket.C {
		log.Println("Ticket tick...")
	}
}

//定时聚合任务
func TickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30
	passengers := make([]string, 0, maxPassenger)
	for {
		passenger := GetNewPassenger()
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}
		select {
		case <-ticker.C:
			Launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger {
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}
func main() {
	TicketDemo()
}
