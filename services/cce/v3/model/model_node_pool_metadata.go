package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NodePoolMetadata struct {
	Name string `json:"name"`

	Uid *string `json:"uid,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`
}

func (o NodePoolMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolMetadata struct{}"
	}

	return strings.Join([]string{"NodePoolMetadata", string(data)}, " ")
}
