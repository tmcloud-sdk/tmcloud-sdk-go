package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type NovaLink struct {
	Href string `json:"href"`

	Rel NovaLinkRel `json:"rel"`
}

func (o NovaLink) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaLink struct{}"
	}

	return strings.Join([]string{"NovaLink", string(data)}, " ")
}

type NovaLinkRel struct {
	value string
}

type NovaLinkRelEnum struct {
	SELF        NovaLinkRel
	BOOKMARK    NovaLinkRel
	ALTERNATE   NovaLinkRel
	DESCRIBEDBY NovaLinkRel
}

func GetNovaLinkRelEnum() NovaLinkRelEnum {
	return NovaLinkRelEnum{
		SELF: NovaLinkRel{
			value: "self",
		},
		BOOKMARK: NovaLinkRel{
			value: "bookmark",
		},
		ALTERNATE: NovaLinkRel{
			value: "alternate",
		},
		DESCRIBEDBY: NovaLinkRel{
			value: "describedby",
		},
	}
}

func (c NovaLinkRel) Value() string {
	return c.value
}

func (c NovaLinkRel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NovaLinkRel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
