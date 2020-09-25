package scaleway

// GetSSHHostname returns the IP of the server for SSH
func (d *Driver) GetSSHHostname() (string, error) {
	return d.IPAddress, nil
}
