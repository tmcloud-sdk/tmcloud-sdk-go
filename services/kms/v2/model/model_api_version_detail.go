package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ApiVersionDetail struct {
	Id *string `json:"id,omitempty"`

	Links *[]ApiLink `json:"links,omitempty"`

	Version *string `json:"version,omitempty"`

	Status *string `json:"status,omitempty"`

	Updated *string `json:"updated,omitempty"`

	MinVersion *string `json:"min_version,omitempty"`
}

func (o ApiVersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiVersionDetail struct{}"
	}

	return strings.Join([]string{"ApiVersionDetail", string(data)}, " ")
}
