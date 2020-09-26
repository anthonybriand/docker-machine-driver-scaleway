package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

// Remove a server (a gracefull stop can be performed before)
func (d *Driver) Remove() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	machineState, err := d.GetState()

	if err != nil {
		return err
	}

	switch machineState {
	case state.Starting:
	case state.Stopping:
	case state.Timeout:
		log.Infof("Server is in " + machineState.String() + " state, retrying in 10 seconds...")
		time.Sleep(10 * time.Second)
		return d.Remove()
	case state.Running:
		err := d.Stop()
		if err != nil {
			return err
		}
	}

	log.Infof("Removing server...")
	err = client.DeleteServer(&instance.DeleteServerRequest{
		ServerID: d.ServerID,
		Zone:     d.Zone,
	})

	if err != nil && (!IsScwError(err) || GetErrorStatus(err) != 404) {
		log.Errorf("Server %s remove failed: %s, retrying in 10 seconds...", d.ServerID, err.Error())
		time.Sleep(10 * time.Second)
		return d.Remove()
	}

	for _, volume := range d.VolumesID {
		err := d.deleteVolume(volume)

		if err != nil {
			log.Errorf("An error occured while deleting the volume: " + volume)
		}
	}

	err = d.deleteIP()

	if err != nil {
		return err
	}

	return nil
}
