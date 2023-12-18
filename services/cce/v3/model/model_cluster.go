package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Cluster struct {
	Kind string `json:"kind"`

	ApiVersion string `json:"apiVersion"`

	Metadata *ClusterMetadata `json:"metadata"`

	Spec *ClusterSpec `json:"spec"`

	Status *ClusterStatus `json:"status,omitempty"`
}

func (o Cluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cluster struct{}"
	}

	return strings.Join([]string{"Cluster", string(data)}, " ")
}
