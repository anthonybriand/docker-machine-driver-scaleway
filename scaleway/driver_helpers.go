//+build test current

package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/scw"
	"os"
)

func GetDriver() *Driver {
	newDriver := NewDriver("", "")
	newDriver.BaseDriver.SSHKeyPath = os.TempDir() + "id_rsa"
	newDriver.Organization = os.Getenv("SCALEWAY_ORGANIZATION")
	newDriver.AccessKey = os.Getenv("SCALEWAY_ACCESS_KEY")
	newDriver.SecretKey = os.Getenv("SCALEWAY_SECRET_KEY")
	newDriver.CommercialType = "DEV1-S"
	newDriver.TerminateOnStop = false
	newDriver.Zone = scw.ZoneFrPar1
	newDriver.IPPersistant = false
	newDriver.image = "docker"
	newDriver.stopping = false
	newDriver.created = false
	newDriver.ipv6 = false

	return newDriver
}
