package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteBackupFileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteBackupFileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackupFileResponse struct{}"
	}

	return strings.Join([]string{"DeleteBackupFileResponse", string(data)}, " ")
}
