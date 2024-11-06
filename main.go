package main

import "fmt"
		"errors"
		"log"

func main () {
	date:= Date{} //Create the struct

	err := date.SetYear(2019) //Create Year value in date struct
	if err != nil {
		log.Fatal(err)
	}

	err := date.SetMonth(5) //Create month value in date struct
	if err != nil {
		log.fatal(err)
	}

	err := date.SetDay(27) //Create day value in date struct
	if err != nil {
		log.fatal(err)
	}
	
	fmt.Println(date) //Print the entire struct
}
