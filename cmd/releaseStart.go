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

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	ghub "github.com/repejota/git-hub"
	"github.com/repejota/git-hub/automation"
	"github.com/spf13/cobra"
)

// ReleaseStartCmd represents the release start command
var ReleaseStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a release",
	Long:  `Start working on an release`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		log.SetFlags(0)

		// by default logging is off
		log.SetOutput(ioutil.Discard)

		// --verbose
		// enable logging if verbose mode
		if VerboseFlag {
			log.SetOutput(os.Stdout)
		}

		path := "."

		// Open repository
		repository, err := ghub.OpenRepository(path)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Open repository at %q successfully\n", path)

		// Get the current branch ( check if we are on master )
		currentBranch, err := automation.GetCurrentBranch()
		if err != nil {
			log.Fatal(err)
		}
		if currentBranch != "master" {
			log.Fatalf("Releases must start from 'master' branch and you are on branch %q", currentBranch)
		}
		log.Printf("You are on %q branch\n", currentBranch)

		// Pull the latest changes from origin master
		log.Printf("Pulling latest changes from origin master\n")
		out, err := automation.PullMasterBranch()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)

		// Calculate new version
		nextVersion, err := repository.NextVersion()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Next version is:", nextVersion)

		// Create local release branch
		releaseBranchName := fmt.Sprintf("release/%s", nextVersion)
		out, err = automation.CreateLocalGitBranch(releaseBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Created local branch: %s\n", releaseBranchName)
		log.Println(out)

		// Push local release branch to origin
		out, err = automation.PushLocalBranchToOrigin(releaseBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)

		// Bump nextVersion
		out, err = automation.BumpNextVersion(nextVersion.String())
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)
	},
}
