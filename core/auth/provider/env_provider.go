// Copyright 2022 TM Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package provider

import (
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/auth"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/auth/basic"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/auth/global"
	"github.com/tmcloud-sdk/tmcloud-sdk-go/core/sdkerr"
	"os"
	"strings"
)

const (
	akEnvName            = "TMCLOUD_SDK_AK"
	skEnvName            = "TMCLOUD_SDK_SK"
	projectIdEnvName     = "TMCLOUD_SDK_PROJECT_ID"
	domainIdEnvName      = "TMCLOUD_SDK_DOMAIN_ID"
	securityTokenEnvName = "TMCLOUD_SDK_SECURITY_TOKEN"
	idpIdEnvName         = "TMCLOUD_SDK_IDP_ID"
	idTokenFileEnvName   = "TMCLOUD_SDK_ID_TOKEN_FILE"
)

type EnvCredentialProvider struct {
	credentialType string
}

// NewEnvCredentialProvider return a env credential provider
// Supported credential types: basic, global
func NewEnvCredentialProvider(credentialType string) *EnvCredentialProvider {
	return &EnvCredentialProvider{credentialType: strings.ToLower(credentialType)}
}

// BasicCredentialEnvProvider return a env provider for basic.Credentials
func BasicCredentialEnvProvider() *EnvCredentialProvider {
	return NewEnvCredentialProvider(basicCredentialType)
}

// GlobalCredentialEnvProvider return a env provider for global.Credentials
func GlobalCredentialEnvProvider() *EnvCredentialProvider {
	return NewEnvCredentialProvider(globalCredentialType)
}

// GetCredentials get basic.Credentials or global.Credentials from environment variables
func (p *EnvCredentialProvider) GetCredentials() (auth.ICredential, error) {
	if p.credentialType == "" {
		return nil, sdkerr.NewCredentialsTypeError("credential type is empty")
	}

	if strings.HasPrefix(p.credentialType, basicCredentialType) {
		builder := basic.NewCredentialsBuilder().WithProjectId(os.Getenv(projectIdEnvName))
		err := fillCommonAttrs(builder, getCommonAttrsFromEnv())
		if err != nil {
			return nil, err
		}
		return builder.Build(), nil
	} else if strings.HasPrefix(p.credentialType, globalCredentialType) {
		builder := global.NewCredentialsBuilder().WithDomainId(os.Getenv(domainIdEnvName))
		err := fillCommonAttrs(builder, getCommonAttrsFromEnv())
		if err != nil {
			return nil, err
		}
		return builder.Build(), nil
	}

	return nil, sdkerr.NewCredentialsTypeError("unsupported credential type: " + p.credentialType)
}

func getCommonAttrsFromEnv() commonAttrs {
	return commonAttrs{
		ak:            os.Getenv(akEnvName),
		sk:            os.Getenv(skEnvName),
		securityToken: os.Getenv(securityTokenEnvName),
		idpId:         os.Getenv(idpIdEnvName),
		idTokenFile:   os.Getenv(idTokenFileEnvName),
	}
}
