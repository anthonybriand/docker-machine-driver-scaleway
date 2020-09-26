//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_GetSSHHostname(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	hostname, err := driver.GetSSHHostname()

	assert.Empty(t, hostname, "IPAddress is not empty on driver creation")
	assert.Nil(t, err, "An error occurred while getting ssh hostname %s", err)

	err = driver.Create()

	assert.Nil(t, err, "Error when creating machine %s", err)

	err = driver.Start()

	assert.Nil(t, err, "Error when starting machine %s %s", driver.ServerID, err)

	sshHostname, err := driver.GetSSHHostname()

	assert.NotEmpty(t, sshHostname, "IPAddress is empty after machine creation")
	assert.Nil(t, err, "An error occurred while getting ssh hostname %s", err)

	err = driver.Stop()

	assert.Nil(t, err, "Error when stopping server %s %s", driver.ServerID, err)

	err = driver.Remove()

	assert.Nil(t, err, "Error when removing server %s %s", driver.ServerID, err)
}
