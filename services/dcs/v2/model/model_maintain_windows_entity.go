package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MaintainWindowsEntity struct {
	Seq *int32 `json:"seq,omitempty"`

	Default *bool `json:"default,omitempty"`

	Begin *string `json:"begin,omitempty"`

	End *string `json:"end,omitempty"`
}

func (o MaintainWindowsEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MaintainWindowsEntity struct{}"
	}

	return strings.Join([]string{"MaintainWindowsEntity", string(data)}, " ")
}
