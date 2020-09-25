package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// getClient Instantiate a api client if needed and return it
func (d *Driver) getClient() (*instance.API, error) {
	if d.api == nil {
		if d.Zone == "" {
			d.Zone = scw.ZoneFrPar1
		}

		client, err := scw.NewClient(
			scw.WithDefaultOrganizationID(d.Organization),
			scw.WithAuth(d.AccessKey, d.SecretKey),
		)

		if err != nil {
			return nil, err
		}

		d.api = instance.NewAPI(client)
	}

	return d.api, nil
}
