package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceImageMembers struct {
	Status string `json:"status"`

	CreatedAt string `json:"created_at"`

	UpdatedAt string `json:"updated_at"`

	ImageId string `json:"image_id"`

	MemberId string `json:"member_id"`

	Schema string `json:"schema"`
}

func (o GlanceImageMembers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceImageMembers struct{}"
	}

	return strings.Join([]string{"GlanceImageMembers", string(data)}, " ")
}
