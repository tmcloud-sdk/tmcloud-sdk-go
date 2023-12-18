package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateScalingNotificationResponse struct {
	TopicUrn *string `json:"topic_urn,omitempty"`

	TopicScene *[]string `json:"topic_scene,omitempty"`

	TopicName      *string `json:"topic_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateScalingNotificationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateScalingNotificationResponse struct{}"
	}

	return strings.Join([]string{"CreateScalingNotificationResponse", string(data)}, " ")
}
