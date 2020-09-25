package scaleway

import (
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
)

// createIP create a persistant ip
func (d *Driver) createIP() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	log.Infof("Creating persistant IP...")
	ip, err := client.CreateIP(&instance.CreateIPRequest{
		Zone:    d.Zone,
		Project: &d.Organization,
		Server:  nil,
		Tags:    nil,
	})

	if err != nil {
		return err
	}

	d.IPID = ip.IP.ID
	d.IPAddress = ip.IP.Address.String()

	return nil
}
