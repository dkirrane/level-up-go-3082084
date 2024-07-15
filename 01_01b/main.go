package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	t, error := time.Parse(expectedFormat, target)
	if error != nil {
		log.Fatal("Canout parse supplied date string", target)
	}
	fmt.Println("DOB is", t)
	return t
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	duration := time.Until(target)
	return duration.Hours() / 24
}

func main() {
	bday := flag.String("bday", "2025-04-24", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
