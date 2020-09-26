package scaleway

import (
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

func (d *Driver) deleteIP() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	if d.IPID != "" && d.IPPersistant {
		log.Infof("Deleting IP: %s", d.IPID)
		err = client.DeleteIP(&instance.DeleteIPRequest{
			Zone: d.Zone,
			IP:   d.IPID,
		})

		if err != nil {
			if IsScwError(err) {
				if GetErrorStatus(err) != 404 {
					log.Infof("Delete of IP %s failed: %s, retrying in 10 seconds...", d.IPID, err.Error())
					time.Sleep(10 * time.Second)
					return d.deleteIP()
				}
			}
			log.Errorf("Delete of IP %s failed: %s", d.IPID, err.Error())
			return err
		}
	}

	return nil
}
