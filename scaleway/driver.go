package scaleway

import (
	"github.com/rancher/machine/libmachine/drivers"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// Driver represents the docker driver interface
type Driver struct {
	*drivers.BaseDriver
	ServerID               string
	Organization           string
	IPID                   string
	AccessKey              string
	SecretKey              string
	CommercialType         string
	FallbackCommercialType []string
	RealCommercialType     string
	TerminateOnStop        bool
	Zone                   scw.Zone
	IPPersistant           bool
	VolumesID              []string
	Tags                   []string
	name                   string
	image                  string
	bootscript             string
	volumes                []string
	stopping               bool
	created                bool
	ipv6                   bool
	api                    *instance.API
}
