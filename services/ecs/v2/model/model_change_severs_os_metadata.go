package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeSeversOsMetadata struct {
	UserData *string `json:"user_data,omitempty"`
}

func (o ChangeSeversOsMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSeversOsMetadata struct{}"
	}

	return strings.Join([]string{"ChangeSeversOsMetadata", string(data)}, " ")
}
