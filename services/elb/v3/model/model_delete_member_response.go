package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type DeleteMemberResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMemberResponse struct{}"
	}

	return strings.Join([]string{"DeleteMemberResponse", string(data)}, " ")
}
