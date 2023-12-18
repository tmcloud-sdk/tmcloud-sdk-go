package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AllQuotas struct {
	Resources *[]AllResources `json:"resources,omitempty"`
}

func (o AllQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AllQuotas struct{}"
	}

	return strings.Join([]string{"AllQuotas", string(data)}, " ")
}
