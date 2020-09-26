package scaleway

import (
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/rancher/machine/libmachine/drivers"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultImage = "docker"
)

// GetCreateFlags registers the flags
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_ACCESS_KEY",
			Name:   "scaleway-access-key",
			Usage:  "Scaleway access key",
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_SECRET_KEY",
			Name:   "scaleway-secret-key",
			Usage:  "Scaleway secret key",
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_ORGANIZATION",
			Name:   "scaleway-organization",
			Usage:  "Scaleway organization",
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_NAME",
			Name:   "scaleway-name",
			Usage:  "Assign a name",
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_COMMERCIAL_TYPE",
			Name:   "scaleway-commercial-type",
			Usage:  "Specifies the commercial type",
			Value:  "DEV1-S",
		},
		mcnflag.StringSliceFlag{
			EnvVar: "SCALEWAY_FALLBACK_COMMERCIAL_TYPE",
			Name:   "scaleway-fallback-commercial-type",
			Usage:  "Specifies the fallback commercial type",
			Value:  []string{"DEV1-M", "DEV1-L", "DEV1-XL", "GP1-XS"},
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_ZONE",
			Name:   "scaleway-zone",
			Usage:  "Specifies the location (" + scw.ZoneFrPar1.String() + "," + scw.ZoneFrPar2.String() + "," + scw.ZoneNlAms1.String() + "," + scw.ZonePlWaw1.String() + ")",
			Value:  scw.ZoneFrPar1.String(),
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_IMAGE",
			Name:   "scaleway-image",
			Usage:  "Specifies the image",
			Value:  defaultImage,
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_BOOTSCRIPT",
			Name:   "scaleway-bootscript",
			Usage:  "Specifies the bootscript",
			Value:  "",
		},
		mcnflag.StringSliceFlag{
			EnvVar: "SCALEWAY_VOLUMES",
			Name:   "scaleway-volumes",
			Usage:  "Attach additional volume (size in GB) (e.g., name1:50,name2:100)",
			Value:  nil,
		},
		mcnflag.StringFlag{
			EnvVar: "SCALEWAY_USER",
			Name:   "scaleway-user",
			Usage:  "Specifies SSH user name",
			Value:  drivers.DefaultSSHUser,
		},
		mcnflag.IntFlag{
			EnvVar: "SCALEWAY_PORT",
			Name:   "scaleway-port",
			Usage:  "Specifies SSH port",
			Value:  drivers.DefaultSSHPort,
		},
		mcnflag.BoolFlag{
			EnvVar: "SCALEWAY_DEBUG",
			Name:   "scaleway-debug",
			Usage:  "Enables Scaleway client debugging",
		},
		mcnflag.BoolFlag{
			EnvVar: "SCALEWAY_IPV6",
			Name:   "scaleway-ipv6",
			Usage:  "Enable ipv6",
		},
		mcnflag.BoolFlag{
			EnvVar: "SCALEWAY_IP_PERSISTANT",
			Name:   "scaleway-ip-persistant",
			Usage:  "Enable ip persistence",
		},
		mcnflag.BoolFlag{
			EnvVar: "SCALEWAY_TERMINATE_ON_STOP",
			Name:   "scaleway-terminate-on-stop",
			Usage:  "Stop the server and remove it",
		},
	}
}
