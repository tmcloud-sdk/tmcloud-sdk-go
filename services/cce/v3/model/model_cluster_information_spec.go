package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ClusterInformationSpec struct {
	Description *string `json:"description,omitempty"`

	CustomSan *[]string `json:"customSan,omitempty"`

	ContainerNetwork *ContainerNetworkUpdate `json:"containerNetwork,omitempty"`

	EniNetwork *EniNetworkUpdate `json:"eniNetwork,omitempty"`

	HostNetwork *ClusterInformationSpecHostNetwork `json:"hostNetwork,omitempty"`
}

func (o ClusterInformationSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInformationSpec struct{}"
	}

	return strings.Join([]string{"ClusterInformationSpec", string(data)}, " ")
}
