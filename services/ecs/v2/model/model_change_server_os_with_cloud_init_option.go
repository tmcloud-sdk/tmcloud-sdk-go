package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithCloudInitOption struct {
	Adminpass *string `json:"adminpass,omitempty"`

	Keyname *string `json:"keyname,omitempty"`

	Userid *string `json:"userid,omitempty"`

	Imageid string `json:"imageid"`

	Metadata *ChangeSeversOsMetadata `json:"metadata,omitempty"`

	Mode *string `json:"mode,omitempty"`
}

func (o ChangeServerOsWithCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithCloudInitOption struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithCloudInitOption", string(data)}, " ")
}
