package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type AddMsdtcRequestBody struct {
	Hosts *[]MsdtcHostOption `json:"hosts,omitempty"`
}

func (o AddMsdtcRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddMsdtcRequestBody struct{}"
	}

	return strings.Join([]string{"AddMsdtcRequestBody", string(data)}, " ")
}
