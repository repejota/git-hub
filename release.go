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

// ReleaseStart ...
func ReleaseStart(path string, gitHubToken string) {
	// Open repository
	repository, err := OpenRepository(path, gitHubToken)
	if err != nil {
		fmt.Println(color.RedString("ERROR: %s", err.Error()))
		os.Exit(1)
	}

	// Get the current branch ( check if we are on master )
	currentBranch, err := automation.GetCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}
	if currentBranch != "master" {
		log.Fatalf("Releases must start from 'master' branch and you are on branch %q", currentBranch)
	}
	fmt.Printf("You are on %q branch\n", currentBranch)

	// Pull the latest changes from origin master
	fmt.Printf("Pulling latest changes from origin master\n")
	out, err := automation.PullMasterBranch()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	// Calculate new version
	nextVersion, err := repository.NextVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Next version is:", nextVersion)

	// Create local release branch
	releaseBranchName := fmt.Sprintf("release/%s", nextVersion)
	out, err = automation.CreateLocalGitBranch(releaseBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created local branch: %s\n", releaseBranchName)
	fmt.Println(out)

	// Push local release branch to origin
	out, err = automation.PushLocalBranchToOrigin(releaseBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)

	// Bump nextVersion
	out, err = automation.BumpNextVersion(nextVersion.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

// ReleaseFinish ...
func ReleaseFinish(path string, gitHubToken string) {
	// Open repository
	repository, err := OpenRepository(path, gitHubToken)
	if err != nil {
		fmt.Println(color.RedString("ERROR: %s", err.Error()))
		os.Exit(1)
	}

	// Get current branch (release branch)
	releaseBranchName, err := automation.GetCurrentBranch()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finishing release", releaseBranchName)

	// Go to master branch
	out, err := automation.GoGitBranch("master")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Checking out master branch")
	fmt.Println(out)

	// Pull and rebase
	out, err = automation.PullAndRebase()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pull and rebase master branch")
	fmt.Println(out)

	// Merge release branch into master
	out, err = automation.MergeBranch(releaseBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Merging %s branch into master\n", releaseBranchName)
	fmt.Println(out)

	// Push changes
	out, err = automation.GitPush()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pushing merge changes into master")
	fmt.Println(out)

	// Create a new version Tag
	currentVersion, err := repository.GetCurrentVersion()
	if err != nil {
		log.Fatal(err)
	}
	out, err = automation.CreateGitTag(currentVersion.String())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Creating a local tag", currentVersion)
	fmt.Println(out)

	// Push tags
	out, err = automation.GitPushTags()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pushing tag", currentVersion)
	fmt.Println(out)

	// Delete remote release branch
	out, err = automation.DeleteRemoteBranch(releaseBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleting remote branch", releaseBranchName)
	fmt.Println(out)

	// Delete local release branch
	out, err = automation.DeleteLocalBranch(releaseBranchName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleting local branch", releaseBranchName)
	fmt.Println(out)
}
