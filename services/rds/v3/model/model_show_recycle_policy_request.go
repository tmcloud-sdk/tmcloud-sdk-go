package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowRecyclePolicyRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowRecyclePolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRecyclePolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowRecyclePolicyRequest", string(data)}, " ")
}
