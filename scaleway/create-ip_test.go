//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_createIP(t *testing.T) {
	driver := GetDriver()
	driver.IPPersistant = true

	err := driver.createIP()

	assert.Nil(t, err, "IP Creation failed: %s", err)
	assert.NotEmpty(t, driver.IPID, "IP ID not affected")
	assert.NotEmpty(t, driver.IPAddress, "IP Address not affected")

	_ = driver.deleteIP()
}
