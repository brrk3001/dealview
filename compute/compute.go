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
	}

	var RecognizedDaysInYear time.Duration = 0
	DaysInYr, error := time.ParseDuration("8760h") //no leap years for now, 8760h in a year

	if error != nil {
		return 0
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
	//fmt.Println(DaysFromRyStartToDealEnd, DaysFromRyStartToDealStart, DaysFromRyEndToDealEnd,DaysFromRyEndToDealStart)

	if DaysFromRyStartToDealEnd < 0 || DaysFromRyEndToDealStart > 0 {
		return 0.0
	}

	if DaysFromRyEndToDealEnd < 0 && DaysFromRyStartToDealStart > 0 {
		return 1.0
	}

	if DaysFromRyStartToDealStart > 0 { //deal ends after RY end but starts within RY
		RecognizedDaysInYear = DaysInYr - DaysFromRyStartToDealStart
	} else { //deal starts before RYstart but ends within RY
		RecognizedDaysInYear = DaysFromRyStartToDealEnd
	}

	return float64(RecognizedDaysInYear / DaysInYr)

}

func main() {

	v2 := "03/01/2007"
	v1 := "02/02/2007"
	v3 := "04/01/2006"
	v4 := "03/31/2007"
	fmt.Println(CompletionPercentage(v1, v2, v3, v4))

}
