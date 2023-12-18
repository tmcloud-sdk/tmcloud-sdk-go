package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type PublicipResult struct {
	Eip *EipResult `json:"eip,omitempty"`
}

func (o PublicipResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipResult struct{}"
	}

	return strings.Join([]string{"PublicipResult", string(data)}, " ")
}
