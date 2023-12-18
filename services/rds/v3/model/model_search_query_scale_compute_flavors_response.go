package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SearchQueryScaleComputeFlavorsResponse struct {
	ComputeFlavorGroups *[]Computes `json:"compute_flavor_groups,omitempty"`
	HttpStatusCode      int         `json:"-"`
}

func (o SearchQueryScaleComputeFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryScaleComputeFlavorsResponse struct{}"
	}

	return strings.Join([]string{"SearchQueryScaleComputeFlavorsResponse", string(data)}, " ")
}
