package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaListKeypairsResult struct {
	Keypair *NovaSimpleKeypair `json:"keypair"`
}

func (o NovaListKeypairsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaListKeypairsResult struct{}"
	}

	return strings.Join([]string{"NovaListKeypairsResult", string(data)}, " ")
}
