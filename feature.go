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
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/repejota/git-hub/automation"
)

// FeatureStart ...
func FeatureStart(repositoryPath string, gitHubToken string, featureTitle string) {
	// Open repository
	_, err := OpenRepository(repositoryPath, gitHubToken)
	if err != nil {
		fmt.Println(color.RedString("ERROR: %s", err.Error()))
		os.Exit(1)
	}

	// Create local issue branch
	featureBranchName := fmt.Sprintf("feature/%s", Slugify(featureTitle))

	out, err := automation.CreateLocalGitBranch(featureBranchName)
	if err != nil {
		fmt.Println(color.RedString("ERROR: %s", err.Error()))
		os.Exit(1)
	}
	fmt.Println("Creating local branch", featureBranchName)
	fmt.Println(out)

	// Push local release branch to origin
	out, err = automation.PushLocalBranchToOrigin(featureBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}
