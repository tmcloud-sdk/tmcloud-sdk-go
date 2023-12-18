package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PersistentVolumeClaimStatus struct {
	AccessModes *[]string `json:"accessModes,omitempty"`

	Capacity *string `json:"capacity,omitempty"`

	Phase *string `json:"phase,omitempty"`
}

func (o PersistentVolumeClaimStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaimStatus struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaimStatus", string(data)}, " ")
}
