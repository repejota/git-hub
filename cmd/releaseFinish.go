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
	"log"

	ghub "github.com/repejota/git-hub"
	"github.com/repejota/git-hub/automation"
	"github.com/spf13/cobra"
)

// ReleaseFinishCmd represents the release finish command
var ReleaseFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish a release",
	Long:  `Finish and publish a release`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		path := "."

		// Open repository
		repository, err := ghub.OpenRepository(path)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Open repository at %q successfully\n", path)

		// Get current branch (release branch)
		releaseBranchName, err := automation.GetCurrentBranch()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Finishing release", releaseBranchName)

		// Go to master branch
		out, err := automation.GoGitBranch("master")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Checking out master branch")
		log.Println(out)

		// Pull and rebase
		out, err = automation.PullAndRebase()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Pull and rebase master branch")
		log.Println(out)

		// Merge release branch into master
		out, err = automation.MergeBranch(releaseBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Merging %s branch into master\n", releaseBranchName)
		log.Println(out)

		// Push changes
		out, err = automation.GitPush()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Pushing merge changes into master")
		log.Println(out)

		// Create a new version Tag
		currentVersion, err := repository.GetCurrentVersion()
		if err != nil {
			log.Fatal(err)
		}
		out, err = automation.CreateGitTag(currentVersion.String())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Creating a local tag", currentVersion)
		log.Println(out)

		// Push tags
		out, err = automation.GitPushTags()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Pushing tag", currentVersion)
		log.Println(out)

		// Delete remote release branch
		out, err = automation.DeleteRemoteBranch(releaseBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deleting remote branch", releaseBranchName)
		log.Println(out)

		// Delete local release branch
		out, err = automation.DeleteLocalBranch(releaseBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Deleting local branch", releaseBranchName)
		log.Println(out)
	},
}
