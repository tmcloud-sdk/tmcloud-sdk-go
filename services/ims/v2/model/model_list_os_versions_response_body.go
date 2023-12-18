package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ListOsVersionsResponseBody struct {
	Platform string `json:"platform"`

	VersionList []OsVersionInfo `json:"version_list"`
}

func (o ListOsVersionsResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOsVersionsResponseBody struct{}"
	}

	return strings.Join([]string{"ListOsVersionsResponseBody", string(data)}, " ")
}
