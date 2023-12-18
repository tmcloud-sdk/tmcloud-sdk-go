package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateServerResult struct {
	TenantId string `json:"tenant_id"`

	Image string `json:"image"`

	AccessIPv4 string `json:"accessIPv4"`

	AccessIPv6 string `json:"accessIPv6"`

	Metadata map[string]string `json:"metadata"`

	Addresses map[string][]UpdateServerAddress `json:"addresses"`

	Created string `json:"created"`

	HostId string `json:"hostId"`

	Flavor *SimpleFlavor `json:"flavor"`

	OSDCFdiskConfig *string `json:"OS-DCF:diskConfig,omitempty"`

	UserId string `json:"user_id"`

	Name string `json:"name"`

	Progress int32 `json:"progress"`

	Links []Link `json:"links"`

	Id string `json:"id"`

	Updated string `json:"updated"`

	Locked *bool `json:"locked,omitempty"`

	Description *string `json:"description,omitempty"`

	Tags []string `json:"tags"`

	Status string `json:"status"`

	OSEXTSRVATTRhostname string `json:"OS-EXT-SRV-ATTR:hostname"`
}

func (o UpdateServerResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerResult struct{}"
	}

	return strings.Join([]string{"UpdateServerResult", string(data)}, " ")
}
