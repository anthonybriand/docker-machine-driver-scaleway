//+build current

package scaleway

import (
	"errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsScwError(t *testing.T) {
	scwError := IsScwError(errors.New("Test"))

	assert.NotNil(t, scwError, "Must not return nil")
	assert.Equal(t, false, scwError, "Is not scaleway error")

	isScwError := IsScwError(scw.ResponseError{
		Message:    "test",
		Type:       "qsdqsd",
		Resource:   "wxcwxc",
		StatusCode: 404,
		Status:     "qsdqsd",
		RawBody:    nil,
	})

	assert.NotNil(t, isScwError, "Must not return nil")
	assert.Equal(t, true, isScwError, "Is a scaleway error")

	isScwError = IsScwError(&scw.ResponseError{
		Message:    "test",
		Type:       "qsdqsd",
		Resource:   "wxcwxc",
		StatusCode: 404,
		Status:     "qsdqsd",
		RawBody:    nil,
	})

	assert.NotNil(t, isScwError, "Must not return nil")
	assert.Equal(t, true, isScwError, "Is a scaleway error pointer")
}

func TestGetErrorStatus(t *testing.T) {
	scwError := GetErrorStatus(errors.New("Test"))

	assert.NotNil(t, scwError, "Must not return nil")
	assert.Equal(t, -1, scwError, "Is not scaleway error so -1 is returned")

	isScwError := GetErrorStatus(scw.ResponseError{
		Message:    "test",
		Type:       "qsdqsd",
		Resource:   "wxcwxc",
		StatusCode: 404,
		Status:     "qsdqsd",
		RawBody:    nil,
	})

	assert.NotNil(t, isScwError, "Must not return nil")
	assert.Equal(t, 404, isScwError, "Is a scaleway error so status code is returned")

	isScwError = GetErrorStatus(&scw.ResponseError{
		Message:    "test",
		Type:       "qsdqsd",
		Resource:   "wxcwxc",
		StatusCode: 404,
		Status:     "qsdqsd",
		RawBody:    nil,
	})

	assert.NotNil(t, isScwError, "Must not return nil")
	assert.Equal(t, 404, isScwError, "Is a scaleway error pointer so status code is returned")
}
