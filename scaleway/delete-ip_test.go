//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_deleteIP(t *testing.T) {
	driver := GetDriver()
	driver.IPPersistant = true

	_ = driver.createIP()

	err := driver.deleteIP()

	assert.Nil(t, err, "Failed to delete IP %s, %s", driver.IPID, err)
}
