//+build test

package scaleway

import (
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_GetCreateFlags(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Driver is nil")

	flags := driver.GetCreateFlags()

	assert.IsType(t, []mcnflag.Flag{}, flags, "Invalid type of flags")

	assert.Len(t, flags, 18, "Invalid flags length")

	assert.Equal(t, flags[0].String(), "scaleway-access-key", "Invalid flag definition at 0")
	assert.Equal(t, flags[1].String(), "scaleway-secret-key", "Invalid flag definition at 1")
	assert.Equal(t, flags[2].String(), "scaleway-organization", "Invalid flag definition at 2")
	assert.Equal(t, flags[3].String(), "scaleway-name", "Invalid flag definition at 3")
	assert.Equal(t, flags[4].String(), "scaleway-commercial-type", "Invalid flag definition at 4")
	assert.Equal(t, flags[5].String(), "scaleway-fallback-commercial-type", "Invalid flag definition at 5")
	assert.Equal(t, flags[6].String(), "scaleway-zone", "Invalid flag definition at 6")
	assert.Equal(t, flags[7].String(), "scaleway-image", "Invalid flag definition at 7")
	assert.Equal(t, flags[8].String(), "scaleway-bootscript", "Invalid flag definition at 8")
	assert.Equal(t, flags[9].String(), "scaleway-volumes", "Invalid flag definition at 9")
	assert.Equal(t, flags[10].String(), "scaleway-user", "Invalid flag definition at 10")
	assert.Equal(t, flags[11].String(), "scaleway-port", "Invalid flag definition at 11")
	assert.Equal(t, flags[12].String(), "scaleway-debug", "Invalid flag definition at 12")
	assert.Equal(t, flags[13].String(), "scaleway-ipv6", "Invalid flag definition at 13")
	assert.Equal(t, flags[14].String(), "scaleway-ip-persistant", "Invalid flag definition at 14")
	assert.Equal(t, flags[15].String(), "scaleway-terminate-on-stop", "Invalid flag definition at 15")
	assert.Equal(t, flags[16].String(), "scaleway-start-on-create", "Invalid flag definition at 16")
	assert.Equal(t, flags[17].String(), "scaleway-tag", "Invalid flag definition at 17")
}
