package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterEndpoints struct {
	Url *string `json:"url,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o ClusterEndpoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterEndpoints struct{}"
	}

	return strings.Join([]string{"ClusterEndpoints", string(data)}, " ")
}
