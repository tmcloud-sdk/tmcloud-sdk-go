package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateKeyRotationIntervalRequestBody struct {
	KeyId string `json:"key_id"`

	RotationInterval int32 `json:"rotation_interval"`

	Sequence *string `json:"sequence,omitempty"`
}

func (o UpdateKeyRotationIntervalRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalRequestBody", string(data)}, " ")
}
