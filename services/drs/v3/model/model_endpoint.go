package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type Endpoint struct {
	DbType *EndpointDbType `json:"db_type,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	Region *string `json:"region,omitempty"`

	InstId *string `json:"inst_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	DbPassword *string `json:"db_password,omitempty"`

	DbPort *int32 `json:"db_port,omitempty"`

	DbUser *string `json:"db_user,omitempty"`

	InstName *string `json:"inst_name,omitempty"`

	Ip *string `json:"ip,omitempty"`

	MongoHaMode *string `json:"mongo_ha_mode,omitempty"`

	SafeMode *int32 `json:"safe_mode,omitempty"`

	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	SslCertName *string `json:"ssl_cert_name,omitempty"`

	SslLink *bool `json:"ssl_link,omitempty"`

	Topic *string `json:"topic,omitempty"`

	ClusterMode *EndpointClusterMode `json:"cluster_mode,omitempty"`
}

func (o Endpoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Endpoint struct{}"
	}

	return strings.Join([]string{"Endpoint", string(data)}, " ")
}

type EndpointDbType struct {
	value string
}

type EndpointDbTypeEnum struct {
	MYSQL       EndpointDbType
	MONGODB     EndpointDbType
	GAUSSDBV5   EndpointDbType
	POSTGRESQL  EndpointDbType
	KAFKA       EndpointDbType
	GAUSSDBV5HA EndpointDbType
}

func GetEndpointDbTypeEnum() EndpointDbTypeEnum {
	return EndpointDbTypeEnum{
		MYSQL: EndpointDbType{
			value: "mysql",
		},
		MONGODB: EndpointDbType{
			value: "mongodb",
		},
		GAUSSDBV5: EndpointDbType{
			value: "gaussdbv5",
		},
		POSTGRESQL: EndpointDbType{
			value: "postgresql",
		},
		KAFKA: EndpointDbType{
			value: "kafka",
		},
		GAUSSDBV5HA: EndpointDbType{
			value: "gaussdbv5ha",
		},
	}
}

func (c EndpointDbType) Value() string {
	return c.value
}

func (c EndpointDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointDbType) UnmarshalJSON(b []byte) error {
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

type EndpointClusterMode struct {
	value string
}

type EndpointClusterModeEnum struct {
	SHARDING4_0 EndpointClusterMode
}

func GetEndpointClusterModeEnum() EndpointClusterModeEnum {
	return EndpointClusterModeEnum{
		SHARDING4_0: EndpointClusterMode{
			value: "Sharding4.0+",
		},
	}
}

func (c EndpointClusterMode) Value() string {
	return c.value
}

func (c EndpointClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointClusterMode) UnmarshalJSON(b []byte) error {
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
