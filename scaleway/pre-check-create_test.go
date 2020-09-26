//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_PreCreateCheck(t *testing.T) {
	driver := GetDriver()

	assert.NotNil(t, driver, "Unable to get driver")

	err := driver.PreCreateCheck()

	assert.Nil(t, err, "Pre-create check failed %s", err)
}
