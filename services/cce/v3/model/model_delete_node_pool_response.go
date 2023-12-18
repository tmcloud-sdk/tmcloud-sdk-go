package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteNodePoolResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodePoolMetadata `json:"metadata,omitempty"`

	Spec *NodePoolSpec `json:"spec,omitempty"`

	Status         *NodePoolStatus `json:"status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o DeleteNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodePoolResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodePoolResponse", string(data)}, " ")
}
