package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"time"
)

// Start start the server
func (d *Driver) Start() error {
	client, err := d.getClient()

	if err != nil {
		return err
	}

	log.Infof("Starting server...")

	machineState, err := d.GetState()

	if err != nil {
		return d.Start()
	}

	if machineState != state.Starting && machineState != state.Running {
		retryInterval := 10 * time.Second
		err = client.ServerActionAndWait(&instance.ServerActionAndWaitRequest{
			ServerID:      d.ServerID,
			Zone:          d.Zone,
			Action:        instance.ServerActionPoweron,
			RetryInterval: &retryInterval,
		})

		if err != nil {
			if err.(*scw.ResponseError).StatusCode == 404 {
				return nil
			} else {
				log.Errorf("Server %s failed to start: %s, retrying in 10 seconds...", d.ServerID, err.Error())
				time.Sleep(10 * time.Second)
				return d.Start()
			}
		}
	}

	serverResponse, err := client.GetServer(&instance.GetServerRequest{
		Zone:     d.Zone,
		ServerID: d.ServerID,
	})

	if err != nil {
		if err.(*scw.ResponseError).StatusCode != 404 {
			return d.Start()
		}
		return nil
	}

	d.IPAddress = serverResponse.Server.PublicIP.Address.String()
	d.IPID = serverResponse.Server.PublicIP.ID

	return nil
}
