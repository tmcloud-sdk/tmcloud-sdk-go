package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowVersionsRequest struct {
}

func (o ShowVersionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVersionsRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionsRequest", string(data)}, " ")
}
