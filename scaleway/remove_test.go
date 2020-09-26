//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_Remove(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	err := driver.Create()

	assert.Nil(t, err, "Failed to create server %s", err)

	err = driver.Remove()

	assert.Nil(t, err, "Failed to remove server %s %s", driver.ServerID, err)
}
