package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateClusterResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *ClusterMetadata `json:"metadata,omitempty"`

	Spec *ClusterSpec `json:"spec,omitempty"`

	Status         *ClusterStatus `json:"status,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o UpdateClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterResponse", string(data)}, " ")
}
