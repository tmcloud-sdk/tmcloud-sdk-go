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

func TestNewRegion(t *testing.T) {
	reg := NewRegion("region-id-1", "https://service1.region-id-1.com")
	assert.NotNil(t, reg)
	assert.Equal(t, "region-id-1", reg.Id)
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, reg.Endpoints)
}

func TestNewRegion2(t *testing.T) {
	reg := NewRegion("region-id-2", "https://service2.region-id-2.com", "https://service2.region-id-2.cn")
	assert.NotNil(t, reg)
	assert.Equal(t, "region-id-2", reg.Id)
	assert.Equal(t, []string{"https://service2.region-id-2.com", "https://service2.region-id-2.cn"}, reg.Endpoints)
}

func TestRegion_WithEndpointsOverride(t *testing.T) {
	reg := NewRegion("region-id-1", "https://service1.region-id-1.com")
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, reg.Endpoints)

	reg.WithEndpointsOverride([]string{"https://service1.region-id-1.cn"})
	assert.Equal(t, []string{"https://service1.region-id-1.cn"}, reg.Endpoints)
}
