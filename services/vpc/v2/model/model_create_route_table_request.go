package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRouteTableRequest struct {
	Body *CreateRoutetableReqBody `json:"body,omitempty"`
}

func (o CreateRouteTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"CreateRouteTableRequest", string(data)}, " ")
}
