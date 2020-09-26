package scaleway

import (
	"errors"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/state"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

// Stop stop the server
func (d *Driver) Stop() error {
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
		return errors.New("Machine " + d.ServerID + " is in an invalid state for stop: " + machineState.String())
	case state.Starting:
		time.Sleep(10 * time.Second)
		err := d.Stop()

		if err != nil {
			return err
		}
	}

	log.Infof("Sending halt to server %s", d.ServerID)
	sshClient, _ := drivers.GetSSHClientFromDriver(d)

	_, _, _ = sshClient.Start("halt")

	_ = sshClient.Wait()

	log.Infof("Stopping server %s", d.ServerID)
	retryInterval := 10 * time.Second
	err = client.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
		ServerID:      d.ServerID,
		Zone:          d.Zone,
		Action:        instance.ServerActionPoweroff,
		RetryInterval: &retryInterval,
	})

	if err != nil {
		if IsScwError(err) && GetErrorStatus(err) == 404 {
			return nil
		} else {
			log.Errorf("Server %s failed to stop: %s, retrying in 10 seconds...", d.ServerID, err.Error())
			return d.Stop()
		}
	}

	if d.TerminateOnStop {
		err := d.Remove()

		if err != nil {
			return err
		}
	}

	return nil
}
