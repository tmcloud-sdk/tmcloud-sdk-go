package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DownloadBackupFilesReq struct {
	Expiration int32 `json:"expiration"`
}

func (o DownloadBackupFilesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadBackupFilesReq struct{}"
	}

	return strings.Join([]string{"DownloadBackupFilesReq", string(data)}, " ")
}
