package scaleway

import "time"

// WaitForTCPPortOpen calls IsTCPPortOpen in a loop
func WaitForTCPPortOpen(dest string) {
	for {
		if IsTCPPortOpen(dest) {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
