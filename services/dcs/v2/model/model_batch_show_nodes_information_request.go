package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type BatchShowNodesInformationRequest struct {
	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`
}

func (o BatchShowNodesInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchShowNodesInformationRequest struct{}"
	}

	return strings.Join([]string{"BatchShowNodesInformationRequest", string(data)}, " ")
}
