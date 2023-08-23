package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	myTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Exact time: %v\n", myTime)
	fmt.Printf("Current time: %v\n", time.Now())
	os.Exit(0)
}
