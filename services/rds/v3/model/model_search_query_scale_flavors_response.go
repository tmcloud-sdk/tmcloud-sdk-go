package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SearchQueryScaleFlavorsResponse struct {
	ComputeFlavorGroups *[]Computes `json:"compute_flavor_groups,omitempty"`
	HttpStatusCode      int         `json:"-"`
}

func (o SearchQueryScaleFlavorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchQueryScaleFlavorsResponse struct{}"
	}

	return strings.Join([]string{"SearchQueryScaleFlavorsResponse", string(data)}, " ")
}
