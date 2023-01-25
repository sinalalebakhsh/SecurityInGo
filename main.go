package main

import (
	"fmt"

	"github.com/elliotforbes/athena/port"
)


func main() {
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort()
}