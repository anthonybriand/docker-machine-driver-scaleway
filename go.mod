module DockerMachineScalewayDriver

go 1.15

replace github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200916142827-bd33bbf0497b+incompatible

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/docker/docker v0.0.0-00010101000000-000000000000 // indirect
	github.com/docker/machine v0.16.2
	github.com/google/go-cmp v0.5.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rancher/machine v0.13.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.6.0.20200923142616-ae5d7880b199
	github.com/sirupsen/logrus v1.7.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gotest.tools v2.2.0+incompatible // indirect
)
