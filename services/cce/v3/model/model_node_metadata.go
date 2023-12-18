package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodeMetadata struct {
	Name *string `json:"name,omitempty"`

	Uid *string `json:"uid,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o NodeMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeMetadata struct{}"
	}

	return strings.Join([]string{"NodeMetadata", string(data)}, " ")
}
