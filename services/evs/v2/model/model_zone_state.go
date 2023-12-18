package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ZoneState struct {
	Available *bool `json:"available,omitempty"`
}

func (o ZoneState) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ZoneState struct{}"
	}

	return strings.Join([]string{"ZoneState", string(data)}, " ")
}
