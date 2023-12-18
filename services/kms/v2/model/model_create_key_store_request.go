package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateKeyStoreRequest struct {
	Body *CreateKeyStoreRequestBody `json:"body,omitempty"`
}

func (o CreateKeyStoreRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKeyStoreRequest struct{}"
	}

	return strings.Join([]string{"CreateKeyStoreRequest", string(data)}, " ")
}
