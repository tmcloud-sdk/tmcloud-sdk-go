package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type KeyStatusInfo struct {
	KeyId *string `json:"key_id,omitempty"`

	KeyState *string `json:"key_state,omitempty"`
}

func (o KeyStatusInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeyStatusInfo struct{}"
	}

	return strings.Join([]string{"KeyStatusInfo", string(data)}, " ")
}
