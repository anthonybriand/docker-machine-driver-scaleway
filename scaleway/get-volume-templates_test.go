//+build test

package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_getVolumeTemplates(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	driver.volumes = []string{"test-vol:50", "test-vol2:100"}

	template, err := driver.getVolumesTemplate()

	assert.Nil(t, err, "Failed to create volume template %s", err)
	assert.Len(t, template, 2, "Invalid template length")
	assert.Equal(t, "test-vol", template["1"].Name, "Invalid name for volume 1")
	assert.Equal(t, "test-vol2", template["2"].Name, "Invalid name for volume 2")
	assert.Equal(t, scw.Size(50)*scw.GB, template["1"].Size, "Invalid size for volume 1")
	assert.Equal(t, scw.Size(100)*scw.GB, template["2"].Size, "Invalid size for volume 2")
	assert.Equal(t, instance.VolumeVolumeTypeLSSD, template["1"].VolumeType, "Invalid size for volume 1")
	assert.Equal(t, instance.VolumeVolumeTypeLSSD, template["2"].VolumeType, "Invalid size for volume 2")
}
