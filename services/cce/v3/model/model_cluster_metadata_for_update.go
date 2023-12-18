package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterMetadataForUpdate struct {
	Alias *string `json:"alias,omitempty"`
}

func (o ClusterMetadataForUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterMetadataForUpdate struct{}"
	}

	return strings.Join([]string{"ClusterMetadataForUpdate", string(data)}, " ")
}
