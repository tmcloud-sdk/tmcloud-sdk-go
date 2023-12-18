package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterInformationSpecHostNetwork struct {
	SecurityGroup *string `json:"SecurityGroup,omitempty"`
}

func (o ClusterInformationSpecHostNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformationSpecHostNetwork struct{}"
	}

	return strings.Join([]string{"ClusterInformationSpecHostNetwork", string(data)}, " ")
}
