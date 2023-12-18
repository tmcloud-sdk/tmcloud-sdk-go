package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ChangeServerOsWithoutCloudInitRequest struct {
	ServerId string `json:"server_id"`

	Body *ChangeServerOsWithoutCloudInitRequestBody `json:"body,omitempty"`
}

func (o ChangeServerOsWithoutCloudInitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeServerOsWithoutCloudInitRequest struct{}"
	}

	return strings.Join([]string{"ChangeServerOsWithoutCloudInitRequest", string(data)}, " ")
}
