package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type OsVersionInfo struct {
	Platform string `json:"platform"`

	OsVersionKey string `json:"os_version_key"`

	OsVersion string `json:"os_version"`

	OsBit int32 `json:"os_bit"`

	OsType string `json:"os_type"`
}

func (o OsVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsVersionInfo struct{}"
	}

	return strings.Join([]string{"OsVersionInfo", string(data)}, " ")
}
