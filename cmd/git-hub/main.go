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

package main

import (
	ghub "github.com/repejota/git-hub"
	"github.com/repejota/git-hub/cmd"
)

var (
	// Version is the current version number
	Version string
	// Build is the current build id
	Build string
)

func main() {
	cmd.RootCmd.SetVersionTemplate(`{{with .Name}}{{printf "%s " .}}{{end}}{{printf "%s" .Version}}`)
	cmd.RootCmd.Version = ghub.ShowVersionInfo(Version, Build)

	cmd.RootCmd.AddCommand(cmd.InfoCmd)

	cmd.IssueCmd.AddCommand(cmd.IssueListCmd)
	cmd.IssueCmd.AddCommand(cmd.IssueStartCmd)
	cmd.IssueCmd.AddCommand(cmd.IssueNewCmd)
	cmd.RootCmd.AddCommand(cmd.IssueCmd)

	cmd.FeatureCmd.AddCommand(cmd.FeatureStartCmd)
	cmd.RootCmd.AddCommand(cmd.FeatureCmd)

	cmd.ReleaseCmd.AddCommand(cmd.ReleaseStartCmd)
	cmd.ReleaseCmd.AddCommand(cmd.ReleaseFinishCmd)
	cmd.RootCmd.AddCommand(cmd.ReleaseCmd)

	cmd.RootCmd.AddCommand(cmd.VersionCmd)

	cmd.Execute()
}
