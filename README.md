## Status

![Pipelines](https://badges.acidspike.fr/badge/60/pipeline)
![Coverage](https://badges.acidspike.fr/badge/60/coverage)
![Quality Gate](https://badges.acidspike.fr/badge/60/alert)
![Reliability](https://badges.acidspike.fr/badge/60/reliability)
![Security](https://badges.acidspike.fr/badge/60/security)
![Vulnerabilities](https://badges.acidspike.fr/badge/60/vulnerability)
![Bugs](https://badges.acidspike.fr/badge/60/bug)
![Code smells](https://badges.acidspike.fr/badge/60/code_smell)

## Overview

A more reliable 3rd-party driver plugin for Docker machine to manage your containers on the servers of Scaleway

## Setup

### Binary

You can find sources and pre-compiled binaries [here](https://github.com/anthonybriand/docker-machine-driver-scaleway/releases/latest)

```shell
# Download the binary (this example downloads the binary for darwin amd64)
$ curl -sL https://github.com/anthonybriand/docker-machine-driver-scaleway/releases/download/v1.4.0/docker-machine-driver-scaleway_darwin64 -O

# Rename the binary, make it executable and copy the binary in a directory accessible with your $PATH
$ mv docker-machine-driver-scaleway_darwin64 docker-machine-driver-scaleway
$ chmod +x docker-machine-driver-scaleway
$ sudo cp docker-machine-driver-scaleway /usr/local/bin/
```

## Usage

### 1. Get your Scaleway credentials

You can generate your `ACCESS KEY` and your `SECRET KEY` [here](https://console.scaleway.com/project/credentials)

### 2. Scaleway driver helper
```console
$ docker-machine create -d scaleway -h
Usage: docker-machine create [OPTIONS] [arg...]

Create a machine

Description:
   Run 'docker-machine create --driver name --help' to include the create flags for that driver in the help text.

Options:
   
   --driver, -d "virtualbox"                                                                                                    Driver to create machine with. [$MACHINE_DRIVER]
   --engine-env [--engine-env option --engine-env option]                                                                       Specify environment variables to set in the engine
   --engine-insecure-registry [--engine-insecure-registry option --engine-insecure-registry option]                             Specify insecure registries to allow with the created engine
   --engine-install-url "https://get.docker.com"                                                                                Custom URL to use for engine installation [$MACHINE_DOCKER_INSTALL_URL]
   --engine-label [--engine-label option --engine-label option]                                                                 Specify labels for the created engine
   --engine-opt [--engine-opt option --engine-opt option]                                                                       Specify arbitrary flags to include with the created engine in the form flag=value
   --engine-registry-mirror [--engine-registry-mirror option --engine-registry-mirror option]                                   Specify registry mirrors to use [$ENGINE_REGISTRY_MIRROR]
   --engine-storage-driver                                                                                                      Specify a storage driver to use with the engine
   --scaleway-access-key                                                                                                        Scaleway access key [$SCALEWAY_ACCESS_KEY]
   --scaleway-bootscript                                                                                                        Specifies the bootscript [$SCALEWAY_BOOTSCRIPT]
   --scaleway-commercial-type "DEV1-S"                                                                                          Specifies the commercial type [$SCALEWAY_COMMERCIAL_TYPE]
   --scaleway-debug                                                                                                             Enables Scaleway client debugging [$SCALEWAY_DEBUG]
   --scaleway-fallback-commercial-type [--scaleway-fallback-commercial-type option --scaleway-fallback-commercial-type option]  Specifies the fallback commercial type [$SCALEWAY_FALLBACK_COMMERCIAL_TYPE]
   --scaleway-image "docker"                                                                                                    Specifies the image [$SCALEWAY_IMAGE]
   --scaleway-ip-persistant                                                                                                     Enable ip persistence [$SCALEWAY_IP_PERSISTANT]
   --scaleway-ipv6                                                                                                              Enable ipv6 [$SCALEWAY_IPV6]
   --scaleway-kill-on-stop                                                                                                      Kill the machine instead of stop it [$SCALEWAY_KILL_ON_STOP]
   --scaleway-name                                                                                                              Assign a name [$SCALEWAY_NAME]
   --scaleway-organization                                                                                                      Scaleway organization [$SCALEWAY_ORGANIZATION]
   --scaleway-port "22"                                                                                                         Specifies SSH port [$SCALEWAY_PORT]
   --scaleway-secret-key                                                                                                        Scaleway secret key [$SCALEWAY_SECRET_KEY]
   --scaleway-start-on-create                                                                                                   Start the server after it was created [$SCALEWAY_START_ON_CREATE]
   --scaleway-tag [--scaleway-tag option --scaleway-tag option]                                                                 Set tags on machine [$SCALEWAY_TAG]
   --scaleway-terminate-on-stop                                                                                                 Stop the server and remove it [$SCALEWAY_TERMINATE_ON_STOP]
   --scaleway-user "root"                                                                                                       Specifies SSH user name [$SCALEWAY_USER]
   --scaleway-volumes [--scaleway-volumes option --scaleway-volumes option]                                                     Attach additional volume (size in GB) (e.g., --scaleway-volumes=name1:50 --scaleway-volumes=name2:100) [$SCALEWAY_VOLUMES]
   --scaleway-zone "fr-par-1"                                                                                                   Specifies the location (fr-par-1,fr-par-2,nl-ams-1,pl-waw-1) [$SCALEWAY_ZONE]
   --swarm                                                                                                                      Configure Machine to join a Swarm cluster
   --swarm-addr                                                                                                                 addr to advertise for Swarm (default: detect and use the machine IP)
   --swarm-discovery                                                                                                            Discovery service to use with Swarm
   --swarm-experimental                                                                                                         Enable Swarm experimental features
   --swarm-host "tcp://0.0.0.0:3376"                                                                                            ip/socket to listen on for Swarm master
   --swarm-image "swarm:latest"                                                                                                 Specify Docker image to use for Swarm [$MACHINE_SWARM_IMAGE]
   --swarm-join-opt [--swarm-join-opt option --swarm-join-opt option]                                                           Define arbitrary flags for Swarm join
   --swarm-master                                                                                                               Configure Machine to be a Swarm master
   --swarm-opt [--swarm-opt option --swarm-opt option]                                                                          Define arbitrary flags for Swarm master
   --swarm-strategy "spread"                                                                                                    Define a default scheduling strategy for Swarm
   --tls-san [--tls-san option --tls-san option]                                                                                Support extra SANs for TLS certs
```

## Options

|Option Name                                                                       |Description              |Default Value                     |required|
|----------------------------------------------------------------------------------|-------------------------|----------------------------------|--------|
|``--scaleway-access-key`` or ``$SCALEWAY_ACCESS_KEY``                             |Access Key               |none                              |yes     |
|``--scaleway-secret-key`` or ``$SCALEWAY_SECRET_KEY``                             |Secret Key               |none                              |yes     |
|``--scaleway-organization`` or ``$SCALEWAY_ORGANIZATION``                         |Organization UUID        |none                              |yes     |
|``--scaleway-commercial-type`` or ``$SCALEWAY_COMMERCIAL_TYPE``                   |Commercial type          |DEV1-S                            |no      |
|``--scaleway-fallback-commercial-type`` or ``$SCALEWAY_FALLBACK_COMMERCIAL_TYPE`` |Fallback commercial type |DEV1-M, DEV1-L, DEV1-XL, GP1-XS   |no      |
|``--scaleway-name`` or ``$SCALEWAY_NAME``                                         |Server name              |none                              |no      |
|``--scaleway-zone`` or ``$SCALEWAY_ZONE``                                         |Specify the location     |fr-par-1                          |no      |
|``--scaleway-image`` or ``$SCALEWAY_IMAGE``                                       |Server image             |docker                            |no      |
|``--scaleway-volumes`` or ``$SCALEWAY_VOLUMES``                                   |Attach additional volume |none                              |no      |
|``--scaleway-user`` or ``$SCALEWAY_USER``                                         |SSH User                 |root                              |no      |
|``--scaleway-port`` or ``$SCALEWAY_PORT``                                         |SSH port                 |22                                |no      |
|``--scaleway-debug`` or ``$SCALEWAY_DEBUG``                                       |Toggle debugging         |false                             |no      |
|``--scaleway-ipv6`` or ``$SCALEWAY_IPV6``                                         |Enable server IPV6       |false                             |no      |
|``--scaleway-ip-persistant`` or ``$SCALEWAY_IP_PERSISTANT``                       |Create a persistent IPV4 |false                             |no      |
|``--scaleway-terminate-on-stop`` or ``$SCALEWAY_TERMINATE_ON_STOP``               |Delete server when stop  |false                             |no      |
|``--scaleway-start-on-create`` or ``$SCALEWAY_START_ON_CREATE``                   |Start server after create|false                             |no      |
|``--scaleway-tag`` or ``$SCALEWAY_TAG``                                           |Set tags on server       |none                              |no      |
|``--scaleway-kill-on-stop`` or ``$SCALEWAY_KILL_ON_STOP``                         |Kill instead of stop     |false                             |no      |

## Debugging

```console
$ SCALEWAY_DEBUG=1 MACHINE_DEBUG=1 docker-machine ...
```

## Development

Feel free to contribute :smiley::beers:

## Links

- **Scaleway console**: https://cloud.scaleway.com/
- **Scaleway sdk**: https://github.com/scaleway/scaleway-sdk-go
- **Docker Machine**: https://docs.docker.com/machine/
- **Report bugs**: https://github.com/anthonybriand/docker-machine-driver-scaleway/issues
- **Original Scaleway Driver**: https://github.com/scaleway/docker-machine-driver-scaleway

## Donate

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donate_SM.gif)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=W5FA7MHSQLH28)

## License

[MIT](https://github.com/anthonybriand/docker-machine-driver-scaleway/blob/master/LICENSE)
