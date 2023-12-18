package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type QuotaResource struct {
	Type *string `json:"type,omitempty"`

	Min *int32 `json:"min,omitempty"`

	Max *int32 `json:"max,omitempty"`

	Quota *int32 `json:"quota,omitempty"`

	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}
