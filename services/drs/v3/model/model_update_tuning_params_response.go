package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type UpdateTuningParamsResponse struct {
	FullSync *[]TuningParameter `json:"full_sync,omitempty"`

	IncreCapture *[]TuningParameter `json:"incre_capture,omitempty"`

	IncreApply *[]TuningParameter `json:"incre_apply,omitempty"`

	IncreRelay *[]TuningParameter `json:"incre_relay,omitempty"`

	ModifyResult   *string `json:"modify_result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateTuningParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTuningParamsResponse struct{}"
	}

	return strings.Join([]string{"UpdateTuningParamsResponse", string(data)}, " ")
}
