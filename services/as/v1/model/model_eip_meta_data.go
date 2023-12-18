package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type EipMetaData struct {
	MetadataBandwidthShareType *string `json:"metadata_bandwidth_share_type,omitempty"`

	MetadataEipId *string `json:"metadata_eip_id,omitempty"`

	MetadataeipAddress *string `json:"metadataeip_address,omitempty"`
}

func (o EipMetaData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EipMetaData struct{}"
	}

	return strings.Join([]string{"EipMetaData", string(data)}, " ")
}
