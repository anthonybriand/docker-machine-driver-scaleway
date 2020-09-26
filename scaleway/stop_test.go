//+build test

package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_Stop(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	err := driver.Create()

	assert.Nil(t, err, "Failed to create server %s", err)

	err = driver.Start()

	assert.Nil(t, err, "Failed to start server %s %s", driver.ServerID, err)

	serverState, err := driver.GetState()

	assert.Nil(t, err, "Failed to get server %s state %s", driver.ServerID, err)
	assert.Equal(t, state.Running, serverState, "Server %s is in invalid state", driver.ServerID)

	err = driver.Stop()

	assert.Nil(t, err, "Failed to stop server %s %s", driver.ServerID, err)

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get server %s state %s", driver.ServerID, err)
	assert.Equal(t, state.Stopped, serverState, "Server %s is in invalid state", driver.ServerID)

	err = driver.Remove()

	assert.Nil(t, err, "Failed to remove server %s %s", driver.ServerID, err)

}
