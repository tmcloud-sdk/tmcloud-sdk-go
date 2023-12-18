package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AddonMetadata struct {
	Uid *string `json:"uid,omitempty"`

	Name *string `json:"name,omitempty"`

	Alias *string `json:"alias,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

func (o AddonMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonMetadata struct{}"
	}

	return strings.Join([]string{"AddonMetadata", string(data)}, " ")
}
