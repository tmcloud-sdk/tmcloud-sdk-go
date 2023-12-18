package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Clusters struct {
	Name *string `json:"name,omitempty"`

	Cluster *ClusterCert `json:"cluster,omitempty"`
}

func (o Clusters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Clusters struct{}"
	}

	return strings.Join([]string{"Clusters", string(data)}, " ")
}
