package scaleway

import (
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"time"
)

// getProductServersAvailability Fetches all the server type and their availability
func (d *Driver) getProductServersAvailability() (*instance.GetServerTypesAvailabilityResponse, error) {

	client, err := d.getClient()

	if err != nil {
		return nil, err
	}

	log.Infof("Retrieving servers availability...")
	availability, err := client.GetServerTypesAvailability(&instance.GetServerTypesAvailabilityRequest{
		Zone: d.Zone,
	})

	if err != nil {
		log.Errorf("Failed to retrieve servers availability: %s, retrying in 10 seconds...", err.Error())
		time.Sleep(10 * time.Second)
		return d.getProductServersAvailability()
	}

	return availability, nil
}
