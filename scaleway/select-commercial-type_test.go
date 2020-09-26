//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_selectCommercialType(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	err := driver.selectCommercialType()

	assert.Nil(t, err, "Failed to select commercial type")
	assert.Equal(t, "DEV1-S", driver.RealCommercialType, "Invalid commercial type")
}
