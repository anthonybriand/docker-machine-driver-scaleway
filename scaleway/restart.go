package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"time"
)

// Restart restart the server
func (d *Driver) Restart() error {
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
		return d.Restart()
	}

	retryInterval := 10 * time.Second
	err = client.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
		ServerID:      d.ServerID,
		Zone:          d.Zone,
		Action:        instance.ServerActionReboot,
		RetryInterval: &retryInterval,
	})

	if err != nil {
		if err.(*scw.ResponseError).StatusCode != 404 {
			log.Errorf("Server %s reboot failed: %s, retrying in 10 seconds...", d.ServerID, err.Error())
			time.Sleep(10 * time.Second)
			return d.Restart()
		}
	}

	return nil
}
