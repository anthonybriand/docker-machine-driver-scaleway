package scaleway

import (
	"bytes"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"time"
)

// setCloudInit Set cloud init for server
func (d *Driver) setCloudInit(client *instance.API, server *instance.CreateServerResponse, cloudInitConfig string) {
	log.Infof("Setting cloud-init config...")

	err := client.SetServerUserData(&instance.SetServerUserDataRequest{
		Zone:     server.Server.Zone,
		ServerID: server.Server.ID,
		Key:      "cloud-init",
		Content:  bytes.NewBufferString(cloudInitConfig),
	})

	if err != nil {
		if err.(*scw.ResponseError).StatusCode != 404 {
			log.Errorf("Set of cloud-init failed on server %s: %s, retrying in 10 seconds...", server.Server.ID, err.Error())
			time.Sleep(10 * time.Second)
			d.setCloudInit(client, server, cloudInitConfig)
		}
	}
}
