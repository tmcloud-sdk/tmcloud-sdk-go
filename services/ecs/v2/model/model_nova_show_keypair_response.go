package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaShowKeypairResponse struct {
	Keypair        *NovaKeypairDetail `json:"keypair,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NovaShowKeypairResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowKeypairResponse struct{}"
	}

	return strings.Join([]string{"NovaShowKeypairResponse", string(data)}, " ")
}
