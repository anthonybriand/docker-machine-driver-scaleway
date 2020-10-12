package scaleway

import (
	"net"
	"time"
)

// IsTCPPortOpen returns true if a TCP communication with "host:port" can be initialized
func IsTCPPortOpen(dest string) bool {
	conn, err := net.DialTimeout("tcp", dest, time.Duration(2000)*time.Millisecond)
	if err == nil {
		defer conn.Close()
	}
	return err == nil
}
