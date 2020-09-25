package scaleway

import "github.com/rancher/machine/libmachine/drivers"

// NewDriver returns a new driver
func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{},
	}
}
