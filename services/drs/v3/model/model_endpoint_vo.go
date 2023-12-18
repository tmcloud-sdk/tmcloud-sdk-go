package model

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/utils"

	"errors"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/converter"

	"strings"
)

type EndpointVo struct {
	Id *string `json:"id,omitempty"`

	ObjId *string `json:"obj_id,omitempty"`

	InstanceName *string `json:"instance_name,omitempty"`

	DbType *EndpointVoDbType `json:"db_type,omitempty"`

	DbUser *string `json:"db_user,omitempty"`

	DbPassword *string `json:"db_password,omitempty"`

	ManageIp *string `json:"manage_ip,omitempty"`

	TrafficIp *string `json:"traffic_ip,omitempty"`

	DbPort *int32 `json:"db_port,omitempty"`

	Region *string `json:"region,omitempty"`

	CreatedAt *string `json:"created_at,omitempty"`

	UpdatedAt *string `json:"updated_at,omitempty"`

	Ip *string `json:"ip,omitempty"`

	PublicIp *string `json:"public_ip,omitempty"`

	AzCode *string `json:"az_code,omitempty"`

	SecurityGroupId *string `json:"security_group_id,omitempty"`

	SubnetId *string `json:"subnet_id,omitempty"`

	VpcId *string `json:"vpc_id,omitempty"`

	VolumeSize *int64 `json:"volume_size,omitempty"`

	FullTransUserPwd *string `json:"full_trans_user_pwd,omitempty"`

	IncrementTransUserPwd *string `json:"increment_trans_user_pwd,omitempty"`

	SslLink *bool `json:"ssl_link,omitempty"`

	SslCertKey *string `json:"ssl_cert_key,omitempty"`

	SslCertName *string `json:"ssl_cert_name,omitempty"`

	SslCertCheckSum *string `json:"ssl_cert_check_sum,omitempty"`

	SslCertPassword *string `json:"ssl_cert_password,omitempty"`

	DbVersion *string `json:"db_version,omitempty"`

	MongoHaMode *EndpointVoMongoHaMode `json:"mongo_ha_mode,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	ClusterMode *EndpointVoClusterMode `json:"cluster_mode,omitempty"`

	InstanceId *string `json:"instance_id,omitempty"`

	DbName *string `json:"db_name,omitempty"`

	Topic *string `json:"topic,omitempty"`

	SafeMode *int32 `json:"safe_mode,omitempty"`

	KerberosVo *KerberosVo `json:"kerberos_vo,omitempty"`

	MultiWriteDbId *string `json:"multi_write_db_id,omitempty"`
}

func (o EndpointVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointVo struct{}"
	}

	return strings.Join([]string{"EndpointVo", string(data)}, " ")
}

type EndpointVoDbType struct {
	value string
}

type EndpointVoDbTypeEnum struct {
	MYSQL   EndpointVoDbType
	MONGODB EndpointVoDbType
}

func GetEndpointVoDbTypeEnum() EndpointVoDbTypeEnum {
	return EndpointVoDbTypeEnum{
		MYSQL: EndpointVoDbType{
			value: "mysql",
		},
		MONGODB: EndpointVoDbType{
			value: "mongodb",
		},
	}
}

func (c EndpointVoDbType) Value() string {
	return c.value
}

func (c EndpointVoDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoDbType) UnmarshalJSON(b []byte) error {
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

type EndpointVoMongoHaMode struct {
	value string
}

type EndpointVoMongoHaModeEnum struct {
	SHARDING       EndpointVoMongoHaMode
	REPLICA_SET    EndpointVoMongoHaMode
	REPLICA_SINGLE EndpointVoMongoHaMode
}

func GetEndpointVoMongoHaModeEnum() EndpointVoMongoHaModeEnum {
	return EndpointVoMongoHaModeEnum{
		SHARDING: EndpointVoMongoHaMode{
			value: "Sharding",
		},
		REPLICA_SET: EndpointVoMongoHaMode{
			value: "ReplicaSet",
		},
		REPLICA_SINGLE: EndpointVoMongoHaMode{
			value: "ReplicaSingle",
		},
	}
}

func (c EndpointVoMongoHaMode) Value() string {
	return c.value
}

func (c EndpointVoMongoHaMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoMongoHaMode) UnmarshalJSON(b []byte) error {
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

type EndpointVoClusterMode struct {
	value string
}

type EndpointVoClusterModeEnum struct {
	SINGLE         EndpointVoClusterMode
	HA             EndpointVoClusterMode
	GR             EndpointVoClusterMode
	SHARDING       EndpointVoClusterMode
	SHARDING4_0    EndpointVoClusterMode
	REPLICA_SET    EndpointVoClusterMode
	REPLICA        EndpointVoClusterMode
	REPLICA_SINGLE EndpointVoClusterMode
	CLUSTER        EndpointVoClusterMode
	INDEPENDENT    EndpointVoClusterMode
	COMBINED       EndpointVoClusterMode
	DISTRIBUTED    EndpointVoClusterMode
	NO_SHARDING    EndpointVoClusterMode
}

func GetEndpointVoClusterModeEnum() EndpointVoClusterModeEnum {
	return EndpointVoClusterModeEnum{
		SINGLE: EndpointVoClusterMode{
			value: "Single",
		},
		HA: EndpointVoClusterMode{
			value: "Ha",
		},
		GR: EndpointVoClusterMode{
			value: "GR",
		},
		SHARDING: EndpointVoClusterMode{
			value: "Sharding",
		},
		SHARDING4_0: EndpointVoClusterMode{
			value: "Sharding4.0+",
		},
		REPLICA_SET: EndpointVoClusterMode{
			value: "ReplicaSet",
		},
		REPLICA: EndpointVoClusterMode{
			value: "Replica",
		},
		REPLICA_SINGLE: EndpointVoClusterMode{
			value: "ReplicaSingle",
		},
		CLUSTER: EndpointVoClusterMode{
			value: "Cluster",
		},
		INDEPENDENT: EndpointVoClusterMode{
			value: "Independent",
		},
		COMBINED: EndpointVoClusterMode{
			value: "Combined",
		},
		DISTRIBUTED: EndpointVoClusterMode{
			value: "Distributed",
		},
		NO_SHARDING: EndpointVoClusterMode{
			value: "NoSharding",
		},
	}
}

func (c EndpointVoClusterMode) Value() string {
	return c.value
}

func (c EndpointVoClusterMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointVoClusterMode) UnmarshalJSON(b []byte) error {
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
