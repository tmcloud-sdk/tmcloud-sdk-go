package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithoutCloudInitOption struct {
	Adminpass *string `json:"adminpass,omitempty"`

	Keyname *string `json:"keyname,omitempty"`

	Userid *string `json:"userid,omitempty"`

	Imageid string `json:"imageid"`

	Mode *string `json:"mode,omitempty"`
}

func (o ChangeServerOsWithoutCloudInitOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitOption struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitOption", string(data)}, " ")
}
