package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type TestEndPoint struct {
	Id string `json:"id"`

	NetType TestEndPointNetType `json:"net_type"`

	DbType TestEndPointDbType `json:"db_type"`

	Ip string `json:"ip"`

	DbPort *int32 `json:"db_port,omitempty"`

	InstId *string `json:"inst_id,omitempty"`

	DbUser string `json:"db_user"`

	DbPassword string `json:"db_password"`

	SslLink *bool `json:"ssl_link,omitempty"`

	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	SslCertName *string `json:"ssl_cert_name,omitempty"`

	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	EndPointType TestEndPointEndPointType `json:"end_point_type"`

	Region *string `json:"region,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	KafkaSecurityConfig *KafkaSecurity `json:"kafka_security_config,omitempty"`
}

func (o TestEndPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestEndPoint struct{}"
	}

	return strings.Join([]string{"TestEndPoint", string(data)}, " ")
}

type TestEndPointNetType struct {
	value string
}

type TestEndPointNetTypeEnum struct {
	VPN TestEndPointNetType
	VPC TestEndPointNetType
	EIP TestEndPointNetType
}

func GetTestEndPointNetTypeEnum() TestEndPointNetTypeEnum {
	return TestEndPointNetTypeEnum{
		VPN: TestEndPointNetType{
			value: "vpn",
		},
		VPC: TestEndPointNetType{
			value: "vpc",
		},
		EIP: TestEndPointNetType{
			value: "eip",
		},
	}
}

func (c TestEndPointNetType) Value() string {
	return c.value
}

func (c TestEndPointNetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointNetType) UnmarshalJSON(b []byte) error {
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

type TestEndPointDbType struct {
	value string
}

type TestEndPointDbTypeEnum struct {
	MYSQL       TestEndPointDbType
	MONGODB     TestEndPointDbType
	GAUSSDBV5   TestEndPointDbType
	POSTGRESQL  TestEndPointDbType
	KAFKA       TestEndPointDbType
	GAUSSDBV5HA TestEndPointDbType
}

func GetTestEndPointDbTypeEnum() TestEndPointDbTypeEnum {
	return TestEndPointDbTypeEnum{
		MYSQL: TestEndPointDbType{
			value: "mysql",
		},
		MONGODB: TestEndPointDbType{
			value: "mongodb",
		},
		GAUSSDBV5: TestEndPointDbType{
			value: "gaussdbv5",
		},
		POSTGRESQL: TestEndPointDbType{
			value: "postgresql",
		},
		KAFKA: TestEndPointDbType{
			value: "kafka",
		},
		GAUSSDBV5HA: TestEndPointDbType{
			value: "gaussdbv5ha",
		},
	}
}

func (c TestEndPointDbType) Value() string {
	return c.value
}

func (c TestEndPointDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointDbType) UnmarshalJSON(b []byte) error {
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

type TestEndPointEndPointType struct {
	value string
}

type TestEndPointEndPointTypeEnum struct {
	SO TestEndPointEndPointType
	TA TestEndPointEndPointType
}

func GetTestEndPointEndPointTypeEnum() TestEndPointEndPointTypeEnum {
	return TestEndPointEndPointTypeEnum{
		SO: TestEndPointEndPointType{
			value: "so",
		},
		TA: TestEndPointEndPointType{
			value: "ta",
		},
	}
}

func (c TestEndPointEndPointType) Value() string {
	return c.value
}

func (c TestEndPointEndPointType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TestEndPointEndPointType) UnmarshalJSON(b []byte) error {
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
