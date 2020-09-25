package scaleway

import (
	"fmt"
	"github.com/docker/machine/libmachine/drivers"
	"net"
)

// GetURL Return the docker URL
func (d *Driver) GetURL() (string, error) {
	if err := drivers.MustBeRunning(d); err != nil {
		return "", err
	}
	return fmt.Sprintf("tcp://%s", net.JoinHostPort(d.IPAddress, "2376")), nil
}
