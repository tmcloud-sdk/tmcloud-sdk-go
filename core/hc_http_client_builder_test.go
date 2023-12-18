// Copyright 2020 TM Technologies Co.,Ltd.
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

package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHcHttpClientBuilder_Build(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	endpoints := []string{"endpoint"}
	err := os.Setenv("TMCLOUD_SDK_CREDENTIALS_FILE", "/cred")
	assert.Nil(t, err)
	client := NewHcHttpClientBuilder().WithEndpoints(endpoints).Build()
	assert.Nil(t, client.credential)

	err = os.Setenv("TMCLOUD_SDK_AK", "ak")
	assert.Nil(t, err)
	err = os.Setenv("TMCLOUD_SDK_SK", "sk")
	assert.Nil(t, err)
	client = NewHcHttpClientBuilder().WithEndpoints(endpoints).Build()
	assert.NotNil(t, client.credential)
}
