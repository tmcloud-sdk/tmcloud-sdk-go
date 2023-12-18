package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"strings"
)

type RoutetableAssociateReqbody struct {
	Routetable *AsscoiateReq `json:"routetable"`
}

func (o RoutetableAssociateReqbody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutetableAssociateReqbody struct{}"
	}

	return strings.Join([]string{"RoutetableAssociateReqbody", string(data)}, " ")
}
