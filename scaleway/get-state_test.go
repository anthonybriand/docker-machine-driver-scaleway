//+build test

package scaleway

import (
	"github.com/docker/machine/libmachine/state"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDriver_GetState(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	serverState, err := driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.None, serverState, "Invalid state when no server")

	err = driver.Create()

	assert.Nil(t, err, "Error when create server %s", err)

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.Stopped, serverState, "Invalid state when server is just created")

	err = driver.Start()

	assert.Nil(t, err, "Error when starting server %s, %s", driver.ServerID, err)

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.Running, serverState, "Invalid state when server is started")

	err = driver.Stop()

	assert.Nil(t, err, "Error when stopping server %s, %s", driver.ServerID, err)

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.Stopped, serverState, "Invalid state when server is stopped")

	action, err := driver.api.ServerAction(&instance.ServerActionRequest{
		Zone:     driver.Zone,
		ServerID: driver.ServerID,
		Action:   "poweron",
	})

	assert.Nil(t, err, "Failed to start server %s", err)
	assert.NotNil(t, action, "No action returned by server action")
	assert.NotEqual(t, instance.TaskStatusFailure, action.Task.Status, "Action has failed")

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.Starting, serverState, "Invalid state when server is starting")

	retryInterval := 10 * time.Second
	_, err = driver.api.WaitForServer(&instance.WaitForServerRequest{
		ServerID:      driver.ServerID,
		Zone:          driver.Zone,
		RetryInterval: &retryInterval,
	})

	assert.Nil(t, err, "Error occurred while waiting for stable server state %s, %s", driver.ServerID, err)

	serverAction, err := driver.api.ServerAction(&instance.ServerActionRequest{
		Zone:     driver.Zone,
		ServerID: driver.ServerID,
		Action:   "poweroff",
	})

	assert.Nil(t, err, "Failed to start server %s", err)
	assert.NotNil(t, serverAction, "No action returned by server action")
	assert.NotEqual(t, instance.TaskStatusFailure, serverAction.Task.Status, "Action has failed")

	serverState, err = driver.GetState()

	assert.Nil(t, err, "Failed to get state %s", err)
	assert.Equal(t, state.Stopping, serverState, "Invalid state when server is stopping")

	retryInterval = 10 * time.Second
	_, err = driver.api.WaitForServer(&instance.WaitForServerRequest{
		ServerID:      driver.ServerID,
		Zone:          driver.Zone,
		RetryInterval: &retryInterval,
	})

	assert.Nil(t, err, "Error occurred while waiting for stable server state %s, %s", driver.ServerID, err)

	err = driver.Remove()

	assert.Nil(t, err, "Error when removing server %s, %s", driver.ServerID, err)
}
