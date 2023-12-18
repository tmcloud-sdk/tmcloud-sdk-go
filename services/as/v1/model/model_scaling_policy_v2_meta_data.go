package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ScalingPolicyV2MetaData struct {
	MetadataBandwidthShareType *string `json:"metadata_bandwidth_share_type,omitempty"`

	MetadataEipId *string `json:"metadata_eip_id,omitempty"`

	MetadataEipAddress *string `json:"metadata_eip_address,omitempty"`
}

func (o ScalingPolicyV2MetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingPolicyV2MetaData struct{}"
	}

	return strings.Join([]string{"ScalingPolicyV2MetaData", string(data)}, " ")
}
