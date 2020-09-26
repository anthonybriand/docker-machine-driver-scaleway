//+build test

package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_getClient(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	client, err := driver.getClient()

	assert.Nil(t, err, "Failed to get client %s", err)
	assert.NotNil(t, client, "Client is nil")
	assert.IsType(t, instance.API{}, *client, "Client has invalid type")
}
