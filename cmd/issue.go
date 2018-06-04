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
	"log"
	"os"
	"strconv"

	ghub "github.com/repejota/git-hub"
	"github.com/repejota/git-hub/automation"
	"github.com/spf13/cobra"
)

// IssueCmd represents the issue command
var IssueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Get information issues",
	Long:  `Get information about the repository issues`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		os.Exit(0)
	},
}

// IssueListCmd represents the issue list command
var IssueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List issues",
	Long:  `List repository issues`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Open repository
		repository, err := ghub.OpenRepository(".")
		if err != nil {
			log.Fatal(err)
		}

		err = repository.GetRemoteGithubRepository("origin")
		if err != nil {
			log.Fatal(err)
		}

		issues, err := ghub.ListIssuesByRepo(*repository.GitHubRepository.FullName)
		if err != nil {
			log.Fatal(err)
		}
		for _, issue := range issues {
			fmt.Printf("#%d - %s - %s\n", *issue.Number, *issue.Title, *issue.HTMLURL)
		}
	},
}

// IssueStartCmd represents the issue start command
var IssueStartCmd = &cobra.Command{
	Use:   "start [issue number]",
	Short: "Start an issue",
	Long:  `Start working on an issue`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatalf("An issue ID is required")
		}
		issueID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid issue ID")
		}
		log.Printf("Start working on the issue #%d\n", issueID)

		repositoryPath := "."

		repository := &ghub.Repository{}

		err = repository.Git(repositoryPath)
		if err != nil {
			log.Fatal(err)
		}

		err = repository.GetRemoteGithubRepository("origin")
		if err != nil {
			log.Fatal(err)
		}

		// Get User
		user, err := ghub.GetAuthenticatedUser()
		if err != nil {
			log.Fatal(err)
		}

		// Get Issue
		org, repo := ghub.ParseRepositoryFullName(*repository.GitHubRepository.FullName)
		issue, err := ghub.GetIssue(org, repo, issueID)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the issue is open
		if issue.GetState() != "open" {
			log.Printf("Issue #%d %q is %s, you can't work on it.\n", issue.GetNumber(), issue.GetTitle(), issue.GetState())
			os.Exit(1)
		}

		// Assign User to the Issue
		err = ghub.AssignUserToIssue(org, repo, user, issue)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Assigned issue to", user.GetLogin())

		// Create local issue branch
		issueBranchName := fmt.Sprintf("issue/%s", ghub.SlugifyIssue(issue))
		out, err := automation.CreateLocalGitBranch(issueBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created local branch", issueBranchName)
		log.Println(out)

		// Push local release branch to origin
		out, err = automation.PushLocalBranchToOrigin(issueBranchName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(out)
	},
}

// IssueFinishCmd represents the issue finish command
var IssueFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish an issue",
	Long:  `Finish working on an issue`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub issue finish")
	},
}
