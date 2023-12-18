package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteStatus struct {
	PreviousTotal *int32 `json:"previous_total,omitempty"`

	CurrentTotal *int32 `json:"current_total,omitempty"`

	Updated *int32 `json:"updated,omitempty"`

	Added *int32 `json:"added,omitempty"`

	Deleted *int32 `json:"deleted,omitempty"`
}

func (o DeleteStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStatus struct{}"
	}

	return strings.Join([]string{"DeleteStatus", string(data)}, " ")
}
