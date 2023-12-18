package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowPublicKeyRequest struct {
	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o ShowPublicKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPublicKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowPublicKeyRequest", string(data)}, " ")
}
