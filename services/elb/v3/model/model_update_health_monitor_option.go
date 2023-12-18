package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type UpdateHealthMonitorOption struct {
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	Delay *int32 `json:"delay,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	ExpectedCodes *string `json:"expected_codes,omitempty"`

	HttpMethod *UpdateHealthMonitorOptionHttpMethod `json:"http_method,omitempty"`

	MaxRetries *int32 `json:"max_retries,omitempty"`

	MaxRetriesDown *int32 `json:"max_retries_down,omitempty"`

	MonitorPort *int32 `json:"monitor_port,omitempty"`

	Name *string `json:"name,omitempty"`

	Timeout *int32 `json:"timeout,omitempty"`

	UrlPath *string `json:"url_path,omitempty"`

	Type *string `json:"type,omitempty"`
}

func (o UpdateHealthMonitorOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHealthMonitorOption struct{}"
	}

	return strings.Join([]string{"UpdateHealthMonitorOption", string(data)}, " ")
}

type UpdateHealthMonitorOptionHttpMethod struct {
	value string
}

type UpdateHealthMonitorOptionHttpMethodEnum struct {
	GET     UpdateHealthMonitorOptionHttpMethod
	HEAD    UpdateHealthMonitorOptionHttpMethod
	POST    UpdateHealthMonitorOptionHttpMethod
	PUT     UpdateHealthMonitorOptionHttpMethod
	DELETE  UpdateHealthMonitorOptionHttpMethod
	TRACE   UpdateHealthMonitorOptionHttpMethod
	OPTIONS UpdateHealthMonitorOptionHttpMethod
	CONNECT UpdateHealthMonitorOptionHttpMethod
	PATCH   UpdateHealthMonitorOptionHttpMethod
}

func GetUpdateHealthMonitorOptionHttpMethodEnum() UpdateHealthMonitorOptionHttpMethodEnum {
	return UpdateHealthMonitorOptionHttpMethodEnum{
		GET: UpdateHealthMonitorOptionHttpMethod{
			value: "GET",
		},
		HEAD: UpdateHealthMonitorOptionHttpMethod{
			value: "HEAD",
		},
		POST: UpdateHealthMonitorOptionHttpMethod{
			value: "POST",
		},
		PUT: UpdateHealthMonitorOptionHttpMethod{
			value: "PUT",
		},
		DELETE: UpdateHealthMonitorOptionHttpMethod{
			value: "DELETE",
		},
		TRACE: UpdateHealthMonitorOptionHttpMethod{
			value: "TRACE",
		},
		OPTIONS: UpdateHealthMonitorOptionHttpMethod{
			value: "OPTIONS",
		},
		CONNECT: UpdateHealthMonitorOptionHttpMethod{
			value: "CONNECT",
		},
		PATCH: UpdateHealthMonitorOptionHttpMethod{
			value: "PATCH",
		},
	}
}

func (c UpdateHealthMonitorOptionHttpMethod) Value() string {
	return c.value
}

func (c UpdateHealthMonitorOptionHttpMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateHealthMonitorOptionHttpMethod) UnmarshalJSON(b []byte) error {
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
