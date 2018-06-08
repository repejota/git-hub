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

	"github.com/fatih/color"
	ghub "github.com/repejota/git-hub"
	"github.com/spf13/cobra"
)

// IssueListCmd represents the issue list command
var IssueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List issues",
	Long:  `List repository issues`,
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

		// --github-token
		// Get the GitHub Token from env or from flag
		gitHubToken := os.Getenv("GITHUB_TOKEN")
		if GitHubToken != "" {
			gitHubToken = GitHubToken
		}

		log.Println(color.YellowString("GitHub Token: %s", gitHubToken))

		// Open repository
		path := "."
		repo, err := ghub.OpenRepository(path, gitHubToken)
		if err != nil {
			fmt.Println(color.RedString("ERROR: %s", err.Error()))
			os.Exit(1)
		}

		// --repository flag
		repository := *repo.GitHubRepository.FullName
		if Repository != "" {
			repository = Repository
		}

		// List issues by repo
		issues, err := ghub.ListIssuesByRepo(repository)
		if err != nil {
			log.Fatal(err)
		}

		// Render issues by repo
		for _, issue := range issues {
			fmt.Printf("#%d - %s - %s\n", *issue.Number, *issue.Title, *issue.HTMLURL)
		}
	},
}

func init() {
	IssueListCmd.Flags().StringVarP(&Repository, "repository", "r", "", "Repository to get the issues from")
}
