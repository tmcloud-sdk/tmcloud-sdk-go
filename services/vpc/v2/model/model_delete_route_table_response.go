package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteRouteTableResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRouteTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRouteTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteRouteTableResponse", string(data)}, " ")
}
