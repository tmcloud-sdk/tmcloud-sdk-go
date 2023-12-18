package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowKeyRotationStatusResponse struct {
	KeyRotationEnabled *bool `json:"key_rotation_enabled,omitempty"`

	RotationInterval *int32 `json:"rotation_interval,omitempty"`

	LastRotationTime *string `json:"last_rotation_time,omitempty"`

	NumberOfRotations *int32 `json:"number_of_rotations,omitempty"`
	HttpStatusCode    int    `json:"-"`
}

func (o ShowKeyRotationStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKeyRotationStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowKeyRotationStatusResponse", string(data)}, " ")
}
