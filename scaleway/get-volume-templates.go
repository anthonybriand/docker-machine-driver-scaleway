package scaleway

import (
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"strconv"
	"strings"
)

// getVolumesTemplate Return the volume objects for API
func (d *Driver) getVolumesTemplate() (map[string]*instance.VolumeTemplate, error) {
	volumesMap := make(map[string]*instance.VolumeTemplate)

	if d.volumes != nil {
		for i, volume := range d.volumes {
			split := strings.Split(volume, ":")
			parseInt := strconv.FormatInt(int64(i+1), 10)
			sizeGB, err := strconv.ParseInt(split[1], 10, 64)

			if err != nil {
				return nil, err
			}

			volumesMap[parseInt] = &instance.VolumeTemplate{
				Name:       split[0],
				Size:       scw.Size(sizeGB) * scw.GB,
				VolumeType: instance.VolumeVolumeTypeLSSD,
				Project:    &d.Organization,
			}
		}
	}

	return volumesMap, nil
}
