package scaleway

import (
	"bytes"
	"github.com/rancher/machine/libmachine/log"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"strings"
	"time"
)

// setAuthorizedKeys Set authorized keys env
func (d *Driver) setAuthorizedKeys(client *instance.API, server *instance.CreateServerResponse, publicKey []byte) {
	log.Infof("Setting Authorized keys config...")
	err := client.SetServerUserData(&instance.SetServerUserDataRequest{
		Zone:     server.Server.Zone,
		ServerID: server.Server.ID,
		Key:      "ENV",
		Content: bytes.NewBufferString(strings.Join([]string{"AUTHORIZED_KEY",
			strings.Replace(string(publicKey[:len(publicKey)-1]), " ", "_", -1)}, "=")),
	})

	if err != nil {
		if err.(*scw.ResponseError).StatusCode != 404 {
			log.Errorf("Set of authorized keys failed on server %s: %s, retrying in 10 seconds...", server.Server.ID, err.Error())
			time.Sleep(10 * time.Second)
			d.setAuthorizedKeys(client, server, publicKey)
		}
	}
}
