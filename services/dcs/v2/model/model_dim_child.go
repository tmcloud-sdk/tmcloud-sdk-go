package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DimChild struct {
	DimName *string `json:"dim_name,omitempty"`

	DimRoute *string `json:"dim_route,omitempty"`
}

func (o DimChild) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimChild struct{}"
	}

	return strings.Join([]string{"DimChild", string(data)}, " ")
}
