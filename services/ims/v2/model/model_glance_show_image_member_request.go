package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type GlanceShowImageMemberRequest struct {
	ImageId string `json:"image_id"`

	MemberId string `json:"member_id"`
}

func (o GlanceShowImageMemberRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlanceShowImageMemberRequest struct{}"
	}

	return strings.Join([]string{"GlanceShowImageMemberRequest", string(data)}, " ")
}
