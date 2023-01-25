package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct  {
	Port int
	State string
}


func ScanPort(protocol, hostname string, port int) bool  {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}



func InitialScan(hostname string,) []ScanResult  {

}