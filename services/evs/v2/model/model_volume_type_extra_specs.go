package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type VolumeTypeExtraSpecs struct {
	RESKEYavailabilityZones *string `json:"RESKEY:availability_zones,omitempty"`

	AvailabilityZone *string `json:"availability-zone,omitempty"`

	OsVendorExtendedsoldOutAvailabilityZones *string `json:"os-vendor-extended:sold_out_availability_zones,omitempty"`

	VolumeBackendName *string `json:"volume_backend_name,omitempty"`

	HWavailabilityZone *string `json:"HW:availability_zone,omitempty"`
}

func (o VolumeTypeExtraSpecs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeTypeExtraSpecs struct{}"
	}

	return strings.Join([]string{"VolumeTypeExtraSpecs", string(data)}, " ")
}
