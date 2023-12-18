package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type NovaKeypair struct {
	Fingerprint string `json:"fingerprint"`

	Name string `json:"name"`

	PublicKey string `json:"public_key"`

	PrivateKey string `json:"private_key"`

	UserId string `json:"user_id"`

	Type *string `json:"type,omitempty"`
}

func (o NovaKeypair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaKeypair struct{}"
	}

	return strings.Join([]string{"NovaKeypair", string(data)}, " ")
}
