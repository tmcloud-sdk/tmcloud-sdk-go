package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ModifyTuningParamsReq struct {
	FullSync map[string]string `json:"full_sync,omitempty"`

	IncreCapture map[string]string `json:"incre_capture,omitempty"`

	IncreApply map[string]string `json:"incre_apply,omitempty"`

	IncreRelay map[string]string `json:"incre_relay,omitempty"`

	Recovery *bool `json:"recovery,omitempty"`
}

func (o ModifyTuningParamsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTuningParamsReq struct{}"
	}

	return strings.Join([]string{"ModifyTuningParamsReq", string(data)}, " ")
}
