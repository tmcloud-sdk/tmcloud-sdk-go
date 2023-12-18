package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ContainerNetworkUpdate struct {
	Cidrs *[]ContainerCidr `json:"cidrs,omitempty"`
}

func (o ContainerNetworkUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerNetworkUpdate struct{}"
	}

	return strings.Join([]string{"ContainerNetworkUpdate", string(data)}, " ")
}
