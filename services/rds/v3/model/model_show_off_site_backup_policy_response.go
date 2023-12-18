package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type ShowOffSiteBackupPolicyResponse struct {
	PolicyPara     *[]GetOffSiteBackupPolicy `json:"policy_para,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowOffSiteBackupPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOffSiteBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowOffSiteBackupPolicyResponse", string(data)}, " ")
}
