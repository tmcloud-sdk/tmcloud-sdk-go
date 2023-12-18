package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterMetadata struct {
	Name string `json:"name"`

	Uid *string `json:"uid,omitempty"`

	Alias *string `json:"alias,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`

	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	UpdateTimestamp *string `json:"updateTimestamp,omitempty"`
}

func (o ClusterMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterMetadata struct{}"
	}

	return strings.Join([]string{"ClusterMetadata", string(data)}, " ")
}
