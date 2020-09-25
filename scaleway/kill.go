package scaleway

import (
	"errors"
	"github.com/docker/machine/libmachine/state"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"time"
)

// Kill force stop a server
func (d *Driver) Kill() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	machineState, err := d.GetState()

	if err != nil {
		return err
	}

	switch machineState {
	case state.Error:
	case state.Timeout:
	case state.Stopping:
	case state.None:
	case state.Stopped:
	case state.Paused:
	case state.Saved:
		return errors.New("Machine is in an invalid state for kill: " + machineState.String())
	case state.Starting:
		log.Infof("Server is in " + machineState.String() + " state, retrying in 10 seconds...")
		time.Sleep(10 * time.Second)
		err := d.Stop()

		if err != nil {
			return err
		}
	}

	retryInterval := 10 * time.Second
	log.Infof("Killing server...")
	err = client.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
		ServerID:      d.ServerID,
		Zone:          d.Zone,
		Action:        instance.ServerActionPoweroff,
		RetryInterval: &retryInterval,
	})

	if err != nil {
		if err.(*scw.ResponseError).StatusCode == 404 {
			return nil
		} else {
			log.Errorf("Failed to kill server %s: %s, retrying in 10 seconds...", d.ServerID, err.Error())
			time.Sleep(10 * time.Second)
			return d.Kill()
		}
	}

	return nil
}
