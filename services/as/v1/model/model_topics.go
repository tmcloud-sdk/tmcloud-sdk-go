package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type Topics struct {
	TopicUrn *string `json:"topic_urn,omitempty"`

	TopicScene *[]string `json:"topic_scene,omitempty"`

	TopicName *string `json:"topic_name,omitempty"`
}

func (o Topics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Topics struct{}"
	}

	return strings.Join([]string{"Topics", string(data)}, " ")
}
