//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_DriverName(t *testing.T) {
	driver := GetDriver()
	driver.RealCommercialType = driver.CommercialType

	name := driver.DriverName()

	assert.Equal(t, "scaleway(DEV1-S)", name, "Invalid driver name %s", name)
}
