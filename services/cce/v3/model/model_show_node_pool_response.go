package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowNodePoolResponse struct {
	Kind *string `json:"kind,omitempty"`

	ApiVersion *string `json:"apiVersion,omitempty"`

	Metadata *NodePoolMetadata `json:"metadata,omitempty"`

	Spec *NodePoolSpec `json:"spec,omitempty"`

	Status         *NodePoolStatus `json:"status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowNodePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodePoolResponse struct{}"
	}

	return strings.Join([]string{"ShowNodePoolResponse", string(data)}, " ")
}
