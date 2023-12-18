package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateMemberOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Name *string `json:"name,omitempty"`

	Weight *int32 `json:"weight,omitempty"`
}

func (o UpdateMemberOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMemberOption struct{}"
	}

	return strings.Join([]string{"UpdateMemberOption", string(data)}, " ")
}
