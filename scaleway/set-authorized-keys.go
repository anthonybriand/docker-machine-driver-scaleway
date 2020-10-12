package scaleway

import (
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"strings"
	"time"
)

// setAuthorizedKeys Set authorized keys env
func (d *Driver) setAuthorizedKeys(client *instance.API, server *instance.CreateServerResponse, publicKey []byte) {
	log.Infof("Setting Authorized keys config...")
	tags := append([]string{strings.Join([]string{"AUTHORIZED_KEY",
		strings.Replace(string(publicKey[:len(publicKey)-1]), " ", "_", -1)}, "=")}, d.Tags...)
	_, err := client.UpdateServer(&instance.UpdateServerRequest{
		Zone:     server.Server.Zone,
		ServerID: server.Server.ID,
		Tags:     &tags,
	})

	if err != nil {
		if !IsScwError(err) || GetErrorStatus(err) != 404 {
			log.Errorf("Set of authorized keys failed on server %s: %s, retrying in 10 seconds...", server.Server.ID, err.Error())
			time.Sleep(10 * time.Second)
			d.setAuthorizedKeys(client, server, publicKey)
		}
	}
}
