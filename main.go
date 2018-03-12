package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	var date string
	if len(os.Args) == 2 {
		date = os.Args[1]
	} else {
		date = "3.7.2018"
	}
	deadline, _ := time.Parse("2.1.2006", date)
	until := time.Until(deadline)
	days := math.Ceil(until.Minutes() / 1440)
	fmt.Println("You've got", days, "days to go.")
}
