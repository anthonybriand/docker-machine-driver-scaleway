package scaleway

import (
	"github.com/docker/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

func (d *Driver) deleteVolume(volume string) error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	log.Infof("Deleting Volume: %s", volume)

	err = client.DeleteVolume(&instance.DeleteVolumeRequest{
		Zone:     d.Zone,
		VolumeID: volume,
	})

	if err != nil {
		if !IsScwError(err) || GetErrorStatus(err) != 404 {
			log.Errorf("Delete of volume %s failed: %s, retrying in 10 seconds...", volume, err.Error())
			time.Sleep(10 * time.Second)
			return d.deleteVolume(volume)
		}
	}

	return nil
}
