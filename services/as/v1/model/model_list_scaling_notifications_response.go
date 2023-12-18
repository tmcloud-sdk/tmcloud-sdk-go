package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListScalingNotificationsResponse struct {
	Topics         *[]Topics `json:"topics,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListScalingNotificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScalingNotificationsResponse struct{}"
	}

	return strings.Join([]string{"ListScalingNotificationsResponse", string(data)}, " ")
}
