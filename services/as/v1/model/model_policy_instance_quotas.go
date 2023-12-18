package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PolicyInstanceQuotas struct {
	Resources *[]PolicyInstanceResources `json:"resources,omitempty"`
}

func (o PolicyInstanceQuotas) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyInstanceQuotas struct{}"
	}

	return strings.Join([]string{"PolicyInstanceQuotas", string(data)}, " ")
}
