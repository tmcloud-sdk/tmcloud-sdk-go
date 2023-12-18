package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowServerAutoRecoveryResponse struct {
	SupportAutoRecovery *string `json:"support_auto_recovery,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowServerAutoRecoveryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerAutoRecoveryResponse struct{}"
	}

	return strings.Join([]string{"ShowServerAutoRecoveryResponse", string(data)}, " ")
}
