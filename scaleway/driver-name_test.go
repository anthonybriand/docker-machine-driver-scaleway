//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_DriverName(t *testing.T) {
	driver := GetDriver()
	driver.RealCommercialType = driver.CommercialType
	driver.ServerName = "test-server"

	name := driver.DriverName()

	assert.Equal(t, "scaleway(DEV1-S,test-server)", name, "Invalid driver name %s", name)
}
