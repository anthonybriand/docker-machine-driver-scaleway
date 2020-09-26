//+build test

package scaleway

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDriver_Create(t *testing.T) {
	driver := GetDriver()

	err := driver.Create()

	assert.Nil(t, err, "Create failed: %s", err)

	_ = driver.Remove()
}
