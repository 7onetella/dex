package net

import (
	"fmt"
	"net"
	"time"
)

// IsTCPConnValid is TCP connection valid
func IsTCPConnValid(host, port string) bool {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", host, port), time.Duration(time.Second*1))
	return err == nil
}
