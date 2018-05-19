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
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
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
	Run: func(cmd *cobra.Command, args []string) {
		// Open repository
		gitRepository, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}
		// Get remotes
		gitRemoteList, err := gitRepository.Remotes()
		if err != nil {
			log.Fatal(err)
		}
		gitRemoteURL := gitRemoteList[0].Config().URLs[0]
		// Get org and repo name
		hubOrganizationName, hubRepositoryName, err := parseRemoteURL(gitRemoteURL)
		if err != nil {
			log.Fatal(err)
		}
		// Get repository issues from Github API
		ctx := context.Background()
		client := github.NewClient(nil)
		options := &github.IssueListByRepoOptions{}
		issues, _, err := client.Issues.ListByRepo(ctx, hubOrganizationName, hubRepositoryName, options)
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
	Use:   "start",
	Short: "Start an issue",
	Long:  `Start working on an issue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub list start")
	},
}

// IssueFinishCmd represents the issue finish command
var IssueFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish an issue",
	Long:  `Finish working on an issue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub issue finish")
	},
}
