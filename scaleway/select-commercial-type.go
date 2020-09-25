package scaleway

import (
	"errors"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
)

// selectCommercialType Select the first available server
func (d *Driver) selectCommercialType() error {
	productServersAvailability, err := d.getProductServersAvailability()

	if err != nil {
		return err
	}

	log.Infof("Removing server...")
	productServer := productServersAvailability.Servers[d.CommercialType]
	commercialType := d.CommercialType
	if productServer.Availability != instance.ServerTypesAvailabilityAvailable && productServer.Availability != instance.ServerTypesAvailabilityScarce {
		commercialType = ""
		for _, stype := range d.FallbackCommercialType {
			if productServersAvailability.Servers[stype].Availability == instance.ServerTypesAvailabilityAvailable || productServersAvailability.Servers[stype].Availability == instance.ServerTypesAvailabilityScarce {
				commercialType = stype
				break
			}
		}
	}

	if commercialType == "" {
		d.created = false
		return errors.New("server type and fallback unavailable")
	}

	d.RealCommercialType = commercialType

	return nil
}
