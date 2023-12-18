package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type CreateOnlineMigrationTaskBody struct {
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	VpcId string `json:"vpc_id"`

	SubnetId string `json:"subnet_id"`

	SecurityGroupId string `json:"security_group_id"`
}

func (o CreateOnlineMigrationTaskBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOnlineMigrationTaskBody struct{}"
	}

	return strings.Join([]string{"CreateOnlineMigrationTaskBody", string(data)}, " ")
}
