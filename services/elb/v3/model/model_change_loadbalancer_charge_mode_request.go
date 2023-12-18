package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeLoadbalancerChargeModeRequest struct {
	Body *ChangeLoadbalancerChargeModeRequestBody `json:"body,omitempty"`
}

func (o ChangeLoadbalancerChargeModeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeRequest struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeRequest", string(data)}, " ")
}
