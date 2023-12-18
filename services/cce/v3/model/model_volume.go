package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Volume struct {
	Size int32 `json:"size"`

	Volumetype string `json:"volumetype"`

	ExtendParam map[string]interface{} `json:"extendParam,omitempty"`

	ClusterId *string `json:"cluster_id,omitempty"`

	ClusterType *string `json:"cluster_type,omitempty"`

	Hwpassthrough *bool `json:"hw:passthrough,omitempty"`

	Metadata *VolumeMetadata `json:"metadata,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
