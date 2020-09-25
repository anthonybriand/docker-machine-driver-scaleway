package main

import (
	"DockerMachineScalewayDriver/scaleway"
	"github.com/rancher/machine/libmachine/drivers/plugin"
)

func main() {
	plugin.RegisterDriver(scaleway.NewDriver("", ""))
}
