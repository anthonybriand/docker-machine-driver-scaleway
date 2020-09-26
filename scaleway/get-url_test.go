//+build test

package scaleway

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestDriver_GetURL(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	url, err := driver.GetURL()

	assert.Error(t, err, "Error not provided")
	assert.Empty(t, url, "Url must be empty")

	err = driver.Create()

	assert.Nil(t, err, "Failed to create server %s", err)

	err = driver.Start()

	assert.Nil(t, err, "Failed to start server %s %s", driver.ServerID, err)

	getURL, err := driver.GetURL()

	assert.Nil(t, err, "Error occurred while getting server %s url %s", driver.ServerID, err)
	assert.NotEmpty(t, getURL, "Url is empty")
	assert.NotEmpty(t, driver.IPAddress, "IP Address is empty")
	assert.Equal(t, fmt.Sprintf("tcp://%s", net.JoinHostPort(driver.IPAddress, "2376")), getURL)

	err = driver.Stop()

	assert.Nil(t, err, "Failed to stop server %s %s", driver.ServerID, err)

	err = driver.Remove()

	assert.Nil(t, err, "Failed to remove server %s %s", driver.ServerID, err)
}
