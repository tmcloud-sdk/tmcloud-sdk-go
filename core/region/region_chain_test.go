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

package region

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProviderChain_GetRegion(t *testing.T) {
	chain := DefaultProviderChain("Service1")
	reg := chain.GetRegion("not-exist-1")
	assert.Nil(t, reg)
}

func TestProviderChain_GetRegion2(t *testing.T) {
	chain := DefaultProviderChain("NotExist")
	reg := chain.GetRegion("region-id-1")
	assert.Nil(t, reg)
}

func TestProviderChain_GetRegion3(t *testing.T) {
	err := setRegionsFileEnv()
	assert.Nil(t, err)

	chain := DefaultProviderChain("Service1")
	reg := chain.GetRegion("region-id-1")
	assert.NotNil(t, reg)
	assert.Equal(t, "region-id-1", reg.Id)
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, reg.Endpoints)
}
