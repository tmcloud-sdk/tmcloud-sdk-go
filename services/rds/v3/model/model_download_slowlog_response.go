package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DownloadSlowlogResponse struct {
	List *[]SlowlogDownloadInfo `json:"list,omitempty"`

	Status *string `json:"status,omitempty"`

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o DownloadSlowlogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSlowlogResponse struct{}"
	}

	return strings.Join([]string{"DownloadSlowlogResponse", string(data)}, " ")
}
