package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PersistentVolumeClaimMetadata struct {
	Name string `json:"name"`

	Labels *string `json:"labels,omitempty"`
}

func (o PersistentVolumeClaimMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistentVolumeClaimMetadata struct{}"
	}

	return strings.Join([]string{"PersistentVolumeClaimMetadata", string(data)}, " ")
}
