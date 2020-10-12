package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

// GetState returns the state of the server
func (d *Driver) GetState() (st state.State, err error) {
	client, err := d.getClient()

	if err != nil {
		return state.Error, err
	}

	if d.ServerID != "" {
		log.Debugf("Retrieving server state...")
		server, err := client.GetServer(&instance.GetServerRequest{
			Zone:     d.Zone,
			ServerID: d.ServerID,
		})

		if err != nil {
			if IsScwError(err) && GetErrorStatus(err) == 404 {
				return state.None, nil
			}
			log.Errorf("Failed to retrieve server %s state: %s, retrying in 10 seconds...", d.ServerID, err.Error())
			time.Sleep(10 * time.Second)
			return d.GetState()
		}

		switch server.Server.State {
		case instance.ServerStateRunning:
			return state.Running, nil
		case instance.ServerStateStarting:
			return state.Starting, nil
		case instance.ServerStateStopped:
			return state.Stopped, nil
		case instance.ServerStateStoppedInPlace:
			return state.Stopped, nil
		case instance.ServerStateStopping:
			return state.Stopping, nil
		default:
			return state.None, nil
		}
	}

	return state.None, nil
}
