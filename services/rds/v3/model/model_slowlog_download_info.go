package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SlowlogDownloadInfo struct {
	WorkflowId string `json:"workflow_id"`

	FileName string `json:"file_name"`

	Status string `json:"status"`

	FileSize string `json:"file_size"`

	FileLink string `json:"file_link"`

	CreateAt int64 `json:"create_at"`

	UpdateAt int64 `json:"update_at"`
}

func (o SlowlogDownloadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowlogDownloadInfo struct{}"
	}

	return strings.Join([]string{"SlowlogDownloadInfo", string(data)}, " ")
}
