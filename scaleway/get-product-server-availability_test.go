//+build test

package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_getProductServerAvailability(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	availability, err := driver.getProductServersAvailability()

	assert.Nil(t, err, "Failed to retrieve server availability %s", err)

	assert.NotNil(t, availability, "Availability is nil")
	assert.IsType(t, instance.GetServerTypesAvailabilityResponse{}, *availability, "Availability is not a valid response")
}
