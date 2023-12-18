package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateClusterEipRequest struct {
	ClusterId string `json:"cluster_id"`

	Body *MasterEipRequest `json:"body,omitempty"`
}

func (o UpdateClusterEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterEipRequest struct{}"
	}

	return strings.Join([]string{"UpdateClusterEipRequest", string(data)}, " ")
}
