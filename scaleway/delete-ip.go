package scaleway

import (
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"time"
)

func (d *Driver) deleteIP() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	log.Infof("Deleting IP: %s", d.IPID)

	err = client.DeleteIP(&instance.DeleteIPRequest{
		Zone: d.Zone,
		IP:   d.IPID,
	})

	if err != nil && err.(*scw.ResponseError).StatusCode != 404 {
		log.Infof("Delete of IP %s failed: %s, retrying in 10 seconds...", d.IPID, err.Error())
		time.Sleep(10 * time.Second)
		return d.deleteIP()
	}

	return nil
}
