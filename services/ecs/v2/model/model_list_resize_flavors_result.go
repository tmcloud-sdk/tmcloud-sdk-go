package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListResizeFlavorsResult struct {
	Id string `json:"id"`

	Name string `json:"name"`

	Vcpus string `json:"vcpus"`

	Ram int32 `json:"ram"`

	Disk string `json:"disk"`

	Swap string `json:"swap"`

	OSFLVEXTDATAephemeral int32 `json:"OS-FLV-EXT-DATA:ephemeral"`

	OSFLVDISABLEDdisabled bool `json:"OS-FLV-DISABLED:disabled"`

	RxtxFactor float32 `json:"rxtx_factor"`

	RxtxQuota string `json:"rxtx_quota"`

	RxtxCap string `json:"rxtx_cap"`

	OsFlavorAccessisPublic bool `json:"os-flavor-access:is_public"`

	Links []FlavorLink `json:"links"`

	ExtraSpecs *FlavorExtraSpec `json:"extra_specs"`

	InstanceQuota *interface{} `json:"instance_quota"`
}

func (o ListResizeFlavorsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResizeFlavorsResult struct{}"
	}

	return strings.Join([]string{"ListResizeFlavorsResult", string(data)}, " ")
}
