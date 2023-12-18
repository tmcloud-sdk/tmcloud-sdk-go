package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateVpcRouteResponse struct {
	Route          *VpcRoute `json:"route,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateVpcRouteResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpcRouteResponse struct{}"
	}

	return strings.Join([]string{"CreateVpcRouteResponse", string(data)}, " ")
}
