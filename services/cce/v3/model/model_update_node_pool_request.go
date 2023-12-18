package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateNodePoolRequest struct {
	ClusterId string `json:"cluster_id"`

	NodepoolId string `json:"nodepool_id"`

	ErrorStatus *string `json:"errorStatus,omitempty"`

	Body *NodePoolUpdate `json:"body,omitempty"`
}

func (o UpdateNodePoolRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodePoolRequest struct{}"
	}

	return strings.Join([]string{"UpdateNodePoolRequest", string(data)}, " ")
}
