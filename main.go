package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)


func main() {
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort("tcp", "localhost", 80)
	fmt.Printf("Port open: %t\n", open)


	results := port.InitialScan("https://google.com")
	fmt.Println(results)
}