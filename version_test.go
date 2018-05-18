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
package hub_test

import (
	"testing"

	hub "github.com/repejota/git-hub"
)

func TestShowVersionInfo(t *testing.T) {
	expectedOutput := "version 1.2.3 build 91b49a2\n"
	version := "1.2.3"
	build := "91b49a2"
	output := hub.ShowVersionInfo(version, build)
	if output != expectedOutput {
		t.Fatalf("Expected output was %q but got %q", expectedOutput, output)
	}
}