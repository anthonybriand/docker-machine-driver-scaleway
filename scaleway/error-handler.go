package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/scw"
	"reflect"
)

func GetErrorStatus(err interface{}) int {
	if IsScwError(err) && reflect.TypeOf(err).Kind() == reflect.Ptr {
		return err.(*scw.ResponseError).StatusCode
	} else if IsScwError(err) {
		return err.(scw.ResponseError).StatusCode
	}

	return -1
}

func IsScwError(err interface{}) bool {
	if reflect.DeepEqual(reflect.TypeOf(err), reflect.TypeOf(scw.ResponseError{})) {
		return true
	} else if reflect.DeepEqual(reflect.TypeOf(err), reflect.PtrTo(reflect.TypeOf(scw.ResponseError{}))) {
		return true
	}

	return false
}
