package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type MasterEipResponseStatus struct {
	PrivateEndpoint *string `json:"privateEndpoint,omitempty"`

	PublicEndpoint *string `json:"publicEndpoint,omitempty"`
}

func (o MasterEipResponseStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MasterEipResponseStatus struct{}"
	}

	return strings.Join([]string{"MasterEipResponseStatus", string(data)}, " ")
}
