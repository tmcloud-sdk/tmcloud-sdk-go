package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ReinstallSeverMetadata struct {
	UserData *string `json:"user_data,omitempty"`
}

func (o ReinstallSeverMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReinstallSeverMetadata struct{}"
	}

	return strings.Join([]string{"ReinstallSeverMetadata", string(data)}, " ")
}
