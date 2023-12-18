package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateSubnetResponse struct {
	Subnet         *Subnet `json:"subnet,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateSubnetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubnetResponse struct{}"
	}

	return strings.Join([]string{"CreateSubnetResponse", string(data)}, " ")
}
