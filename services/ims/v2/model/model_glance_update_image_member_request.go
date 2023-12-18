package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceUpdateImageMemberRequest struct {
	ImageId string `json:"image_id"`

	MemberId string `json:"member_id"`

	Body *GlanceUpdateImageMemberRequestBody `json:"body,omitempty"`
}

func (o GlanceUpdateImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceUpdateImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceUpdateImageMemberRequest", string(data)}, " ")
}
