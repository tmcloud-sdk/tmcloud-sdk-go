package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteScalingTagInfoResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingTagInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingTagInfoResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingTagInfoResponse", string(data)}, " ")
}
