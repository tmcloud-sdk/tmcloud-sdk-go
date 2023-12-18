package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PolicyInstanceResources struct {
	Type *string `json:"type,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Quota *int32 `json:"quota,omitempty"`

	Max *int32 `json:"max,omitempty"`

	Min *int32 `json:"min,omitempty"`
}

func (o PolicyInstanceResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyInstanceResources struct{}"
	}

	return strings.Join([]string{"PolicyInstanceResources", string(data)}, " ")
}
