package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateClusterEipResponse struct {
	Metadata *Metadata `json:"metadata,omitempty"`

	Spec *MasterEipResponseSpec `json:"spec,omitempty"`

	Status         *MasterEipResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o UpdateClusterEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterEipResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterEipResponse", string(data)}, " ")
}
