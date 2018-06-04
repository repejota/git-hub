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

package ghub_test

import (
	"testing"

	"github.com/repejota/git-hub"
)

func TestSemVerInstance(t *testing.T) {
	_, err := ghub.NewSemVer("1.2.3")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSemVerInstanceStringer(t *testing.T) {
	expectedVersion := "1.2.3"

	version, err := ghub.NewSemVer(expectedVersion)
	if err != nil {
		t.Fatal(err)
	}

	if version.String() != "1.2.3" {
		t.Fatalf("Version expected to be %q but got %q", expectedVersion, version.String())
	}
}
