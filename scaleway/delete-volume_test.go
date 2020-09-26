//+build test

package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_deleteVolume(t *testing.T) {
	driver := GetDriver()
	client, err := driver.getClient()

	assert.Nil(t, err, "Failed to get client %s", err)

	volume, err := client.CreateVolume(&instance.CreateVolumeRequest{
		Zone:       driver.Zone,
		Name:       "test_volume",
		Project:    &driver.Organization,
		VolumeType: instance.VolumeVolumeTypeLSSD,
		Size:       scw.SizePtr(scw.Size(50) * scw.GB),
	})

	assert.Nil(t, err, "Failed to create volume %s", err)

	assert.NotNil(t, volume, "Creation not returned volume")

	assert.NotEmpty(t, volume.Volume.ID, "No ID returned when creating volume")

	err = driver.deleteVolume(volume.Volume.ID)

	assert.Nil(t, err, "Failed to remove volume %s, %s", volume.Volume.ID, err)
}
