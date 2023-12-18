package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ReinstallServerWithCloudInitOption struct {
	Adminpass *string `json:"adminpass,omitempty"`

	Keyname *string `json:"keyname,omitempty"`

	Userid *string `json:"userid,omitempty"`

	Metadata *ReinstallSeverMetadata `json:"metadata,omitempty"`

	Mode *string `json:"mode,omitempty"`
}

func (o ReinstallServerWithCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallServerWithCloudInitOption struct{}"
	}

	return strings.Join([]string{"ReinstallServerWithCloudInitOption", string(data)}, " ")
}
