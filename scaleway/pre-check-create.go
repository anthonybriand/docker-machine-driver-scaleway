package scaleway

// PreCreateCheck Check if server can be instantiated
func (d *Driver) PreCreateCheck() error {
	return d.selectCommercialType()
}
