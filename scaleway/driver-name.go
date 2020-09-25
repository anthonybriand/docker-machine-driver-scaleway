package scaleway

import "fmt"

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	if d.RealCommercialType == "" {
		return "scaleway"
	}
	return fmt.Sprintf("scaleway(%v)", d.RealCommercialType)
}
