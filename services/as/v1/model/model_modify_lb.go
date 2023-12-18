package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyLb struct {
	LbaasListener *LbaasListener `json:"lbaas_listener,omitempty"`

	Listener *string `json:"listener,omitempty"`

	FailedReason *string `json:"failed_reason,omitempty"`

	FailedDetails *string `json:"failed_details,omitempty"`
}

func (o ModifyLb) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyLb struct{}"
	}

	return strings.Join([]string{"ModifyLb", string(data)}, " ")
}
