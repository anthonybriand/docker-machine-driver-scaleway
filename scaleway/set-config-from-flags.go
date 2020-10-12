package scaleway

import (
	"fmt"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/sirupsen/logrus"
	"os"
)

// SetConfigFromFlags sets the flags
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) (err error) {
	if flags.Bool("scaleway-debug") {
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.DebugLevel)
	}

	d.AccessKey, d.SecretKey, d.Organization = flags.String("scaleway-access-key"), flags.String("scaleway-secret-key"), flags.String("scaleway-organization")
	if d.AccessKey == "" || d.SecretKey == "" || d.Organization == "" {
		config, cfgErr := scw.LoadConfig()
		if cfgErr == nil {
			if d.AccessKey == "" {
				d.AccessKey = *config.AccessKey
			}
			if d.Organization == "" {
				d.Organization = *config.DefaultOrganizationID
			}
			if d.SecretKey == "" {
				d.SecretKey = *config.SecretKey
			}
		} else {
			return fmt.Errorf("You must provide organization, access key and secret key")
		}
	}
	d.CommercialType = flags.String("scaleway-commercial-type")
	d.FallbackCommercialType = flags.StringSlice("scaleway-fallback-commercial-type")
	switch flags.String("scaleway-zone") {
	case scw.ZoneFrPar1.String():
		d.Zone = scw.ZoneFrPar1
		break
	case scw.ZoneFrPar2.String():
		d.Zone = scw.ZoneFrPar2
		break
	case scw.ZoneNlAms1.String():
		d.Zone = scw.ZoneNlAms1
		break
	case scw.ZonePlWaw1.String():
		break
	default:
		d.Zone = scw.ZoneFrPar1
	}
	d.name = flags.String("scaleway-name")
	d.image = flags.String("scaleway-image")
	d.bootscript = flags.String("scaleway-bootscript")
	d.volumes = flags.StringSlice("scaleway-volumes")
	d.ipv6 = flags.Bool("scaleway-ipv6")
	d.BaseDriver.SSHUser = flags.String("scaleway-user")
	d.BaseDriver.SSHPort = flags.Int("scaleway-port")
	d.TerminateOnStop = flags.Bool("scaleway-terminate-on-stop")
	d.IPPersistant = flags.Bool("scaleway-ip-persistant")
	d.StartOnCreate = flags.Bool("scaleway-start-on-create")
	return
}
