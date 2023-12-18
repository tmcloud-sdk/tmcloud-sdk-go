package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RegisterServerAutoRecoveryRequest struct {
	ServerId string `json:"server_id"`

	Body *RegisterServerAutoRecoveryRequestBody `json:"body,omitempty"`
}

func (o RegisterServerAutoRecoveryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterServerAutoRecoveryRequest struct{}"
	}

	return strings.Join([]string{"RegisterServerAutoRecoveryRequest", string(data)}, " ")
}
