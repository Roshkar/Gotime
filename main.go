package main

import (
	"fmt"
	"time"
	_ "time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Hello! The time is: %s\n", t.Format("02.01.2006 15.04"))
}
