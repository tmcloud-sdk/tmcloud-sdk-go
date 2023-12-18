package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type ListInstancesRequest struct {
	ContentType *string `json:"Content-Type,omitempty"`

	XLanguage *ListInstancesRequestXLanguage `json:"X-Language,omitempty"`

	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Type *ListInstancesRequestType `json:"type,omitempty"`

	DatastoreType *ListInstancesRequestDatastoreType `json:"datastore_type,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	Limit *int32 `json:"limit,omitempty"`

	Tags *string `json:"tags,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestXLanguage struct {
	value string
}

type ListInstancesRequestXLanguageEnum struct {
	ZH_CN ListInstancesRequestXLanguage
	EN_US ListInstancesRequestXLanguage
}

func GetListInstancesRequestXLanguageEnum() ListInstancesRequestXLanguageEnum {
	return ListInstancesRequestXLanguageEnum{
		ZH_CN: ListInstancesRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ListInstancesRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ListInstancesRequestXLanguage) Value() string {
	return c.value
}

func (c ListInstancesRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestType struct {
	value string
}

type ListInstancesRequestTypeEnum struct {
	SINGLE  ListInstancesRequestType
	HA      ListInstancesRequestType
	REPLICA ListInstancesRequestType
}

func GetListInstancesRequestTypeEnum() ListInstancesRequestTypeEnum {
	return ListInstancesRequestTypeEnum{
		SINGLE: ListInstancesRequestType{
			value: "Single",
		},
		HA: ListInstancesRequestType{
			value: "Ha",
		},
		REPLICA: ListInstancesRequestType{
			value: "Replica",
		},
	}
}

func (c ListInstancesRequestType) Value() string {
	return c.value
}

func (c ListInstancesRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestType) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestDatastoreType struct {
	value string
}

type ListInstancesRequestDatastoreTypeEnum struct {
	MY_SQL      ListInstancesRequestDatastoreType
	POSTGRE_SQL ListInstancesRequestDatastoreType
	SQL_SERVER  ListInstancesRequestDatastoreType
}

func GetListInstancesRequestDatastoreTypeEnum() ListInstancesRequestDatastoreTypeEnum {
	return ListInstancesRequestDatastoreTypeEnum{
		MY_SQL: ListInstancesRequestDatastoreType{
			value: "MySQL",
		},
		POSTGRE_SQL: ListInstancesRequestDatastoreType{
			value: "PostgreSQL",
		},
		SQL_SERVER: ListInstancesRequestDatastoreType{
			value: "SQLServer",
		},
	}
}

func (c ListInstancesRequestDatastoreType) Value() string {
	return c.value
}

func (c ListInstancesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
