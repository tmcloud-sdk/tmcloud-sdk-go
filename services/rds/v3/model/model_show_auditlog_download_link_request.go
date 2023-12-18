package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowAuditlogDownloadLinkRequest struct {
	XLanguage *string `json:"X-Language,omitempty"`

	InstanceId string `json:"instance_id"`

	Body *GenerateAuditlogDownloadLinkRequest `json:"body,omitempty"`
}

func (o ShowAuditlogDownloadLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditlogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditlogDownloadLinkRequest", string(data)}, " ")
}
