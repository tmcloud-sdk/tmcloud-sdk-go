package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PersistentVolumeClaim struct {
	ApiVersion string `json:"apiVersion"`

	Kind string `json:"kind"`

	Metadata *PersistentVolumeClaimMetadata `json:"metadata"`

	Spec *PersistentVolumeClaimSpec `json:"spec"`

	Status *PersistentVolumeClaimStatus `json:"status,omitempty"`
}

func (o PersistentVolumeClaim) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaim struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaim", string(data)}, " ")
}
