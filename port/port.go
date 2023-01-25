package port

import "strconv"

func ScanPort(protocol, hostname string, port int) bool  {
	address := hostname + ":" + strconv.Itoa(port)
}