package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateRedislogDownloadLinkRequest struct {
	InstanceId string `json:"instance_id"`

	Id string `json:"id"`
}

func (o CreateRedislogDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRedislogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"CreateRedislogDownloadLinkRequest", string(data)}, " ")
}
