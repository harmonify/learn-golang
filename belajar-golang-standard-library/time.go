package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.UTC())

	utc := time.Date(2009, time.November, 10, 17, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	valueDateTime, err := time.Parse(time.DateTime, "2024-04-07 19:00:00")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("DateTime:", valueDateTime)
	}

	valueRFC3339, err := time.Parse(time.RFC3339, "2024-04-07T19:00:00+07:00")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("RFC3339:", valueRFC3339)
		fmt.Println("RFC3339 year:", valueRFC3339.Year())
	}

	invalidValue, err3 := time.Parse(time.RFC3339, "ASAL_ASAL")
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("RFC3339:", invalidValue)
	}
}
