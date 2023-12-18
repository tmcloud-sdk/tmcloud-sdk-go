package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type SetOffSiteBackupPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetOffSiteBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetOffSiteBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetOffSiteBackupPolicyResponse", string(data)}, " ")
}
