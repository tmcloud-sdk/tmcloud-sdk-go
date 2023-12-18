package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Resources struct {
	Unit *string `json:"unit,omitempty"`

	Min *int32 `json:"min,omitempty"`

	Max *int32 `json:"max,omitempty"`

	Quota *int32 `json:"quota,omitempty"`

	Used *int32 `json:"used,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}
