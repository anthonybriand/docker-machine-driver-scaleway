package scaleway

import (
	"fmt"
	"github.com/rancher/machine/libmachine/log"
	"github.com/rancher/machine/libmachine/ssh"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"io/ioutil"
)

// Create create and configures a scaleway server
func (d *Driver) Create() error {

	var publicKey []byte

	log.Infof("Creating SSH key...")
	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return err
	}
	publicKey, err := ioutil.ReadFile(d.GetSSHKeyPath() + ".pub")
	if err != nil {
		return err
	}

	log.Infof("Creating server...")
	err = d.selectCommercialType()

	if err != nil {
		return err
	}

	bootType := instance.BootTypeBootscript

	cloudInitConfig := fmt.Sprintf(`#cloud-config

# Some images do not have sudo installed by default.
packages:
    - sudo

# Add root to sudoers.
users:
  - name: root
    ssh-authorized-keys: [%s]
    sudo: ['ALL=(ALL) NOPASSWD:ALL']
    groups: sudo
`, publicKey)

	client, err := d.getClient()

	if err != nil {
		return err
	}

	volumesTemplate, err := d.getVolumesTemplate()

	if err != nil {
		return err
	}

	if d.IPPersistant {
		err := d.createIP()
		if err != nil {
			return err
		}
	}

	createServerRequest := &instance.CreateServerRequest{
		Zone:              d.Zone,
		Name:              d.name,
		DynamicIPRequired: scw.BoolPtr(!d.IPPersistant),
		CommercialType:    d.RealCommercialType,
		Image:             d.image,
		EnableIPv6:        d.ipv6,
		BootType:          &bootType,
		Project:           &d.Organization,
		Tags:              d.Tags,
	}

	if len(volumesTemplate) > 0 {
		createServerRequest.Volumes = volumesTemplate
	}

	if d.IPID != "" {
		createServerRequest.PublicIP = &d.IPID
	}

	if d.bootscript != "" {
		createServerRequest.Bootscript = &d.bootscript
	} else {
		bootType = instance.BootTypeLocal
		createServerRequest.BootType = &bootType
	}

	server, err := client.CreateServer(createServerRequest)

	if err != nil {
		_ = d.deleteIP()
		return err
	}

	d.ServerID = server.Server.ID
	d.MachineName = server.Server.Name

	d.setAuthorizedKeys(client, server, publicKey)

	d.setCloudInit(client, server, cloudInitConfig)

	for _, volume := range server.Server.Volumes {
		d.VolumesID = append(d.VolumesID, volume.ID)
	}

	d.created = true

	return nil
}
