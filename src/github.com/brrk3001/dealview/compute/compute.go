package main

import (
	"fmt"
	"time"
)

const (
	// See http://golang.org/pkg/time/#Parse
	TimeFormat = "01/02/2006"
)

func CompletionPercentage(DateStart, DateEnd, RevYearStart, RevYearEnd string) float64 {

	start, error1 := time.Parse(TimeFormat, DateStart)
	end, error2 := time.Parse(TimeFormat, DateEnd)
	rystart, error3 := time.Parse(TimeFormat, RevYearStart)
	ryend, error4 := time.Parse(TimeFormat, RevYearEnd)

	if error1 != nil || error2 != nil || error3 != nil || error4 != nil {

		fmt.Println(error1, error2, error3, error4)
		panic(fmt.Sprintf("Exiting"))
	}

	var RecognizedDaysInYear time.Duration = 0
	DaysInYr, error := time.ParseDuration("8760h") //no leap years for now, 8760h in a year

	if error != nil {
		fmt.Println(error)
		panic(fmt.Sprintf("Exiting"))
	}
	DealDuration := end.Sub(start)
	fmt.Println(end, start, DealDuration)

	if DealDuration < 0 {
		panic(fmt.Sprintf("End date is earlier than start date"))
	}

	DaysFromRyStartToDealEnd := end.Sub(rystart)
	DaysFromRyStartToDealStart := start.Sub(rystart)
	DaysFromRyEndToDealEnd := end.Sub(ryend)
	DaysFromRyEndToDealStart := start.Sub(ryend)

	fmt.Println(DaysFromRyStartToDealEnd, DaysFromRyStartToDealStart, DaysFromRyEndToDealEnd, DaysFromRyEndToDealStart)

	if DaysFromRyStartToDealEnd < 0 || DaysFromRyEndToDealStart > 0 {
		return 0.0
	}

	if DaysFromRyEndToDealEnd <= 0 && DaysFromRyStartToDealStart >= 0 { //starts, ends completely within year
		return 1.0
	}

	if DaysFromRyEndToDealEnd > 0 { //deal ends after RY end but starts within or before RY

		if DaysFromRyStartToDealStart > 0 { //Starts after RY but ends after RYend
			RecognizedDaysInYear = DaysInYr - DaysFromRyStartToDealStart
		} else {
			RecognizedDaysInYear = DaysInYr
		}
	} else { //deal starts before RYstart but ends within  or at RY
		RecognizedDaysInYear = DaysFromRyStartToDealEnd
	}

	return float64(RecognizedDaysInYear.Hours() / DealDuration.Hours())
}

func TestCompletionPercentage() {

	end := "08/02/2007"
	start := "04/02/2006"
	rystart := "04/01/2006"
	ryend := "03/31/2007"
	fmt.Println(CompletionPercentage(start, end, rystart, ryend))

}

func main() {
	TestCompletionPercentage()

}
