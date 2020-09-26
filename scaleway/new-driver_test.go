//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDriver(t *testing.T) {
	driver := NewDriver("", "")

	assert.NotNil(t, driver, "Driver is nil")
	assert.IsType(t, Driver{}, *driver, "Driver is not of type Driver")
}
