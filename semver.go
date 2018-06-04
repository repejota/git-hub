// Copyright 2018 Raül Pérez, repejota@gmail.com. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package ghub

import (
	"fmt"
	"strconv"
	"strings"
)

// SemVer ...
type SemVer struct {
	Major int
	Minor int
	Patch int
}

// NewSemVer ...
func NewSemVer(version string) (*SemVer, error) {
	semver := &SemVer{}
	dataParts := strings.Split(string(version), ".")
	if len(dataParts) < 2 {
		return nil, fmt.Errorf("ERROR calculating next version, invalid VERSION contents: %s", string(version))
	}
	part, err := strconv.Atoi(dataParts[0])
	if err != nil {
		return nil, fmt.Errorf("ERROR calculating next version, invalid Major version:%q", dataParts)
	}
	semver.Major = part
	part, err = strconv.Atoi(dataParts[1])
	if err != nil {
		return nil, fmt.Errorf("ERROR calculating next version, invalid Minor version: %q", dataParts)
	}
	semver.Minor = part
	part, err = strconv.Atoi(dataParts[2])
	if err != nil {
		return nil, fmt.Errorf("ERROR calculating next version, invalid Patch version: %q", dataParts)
	}
	semver.Patch = part
	return semver, nil
}

func (v *SemVer) String() string {
	strSemVer := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	return strSemVer
}
