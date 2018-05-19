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
	"strings"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	git "gopkg.in/src-d/go-git.v4"
)

// InfoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get information about the repository",
	Long:  `Get information about the repository and its github project`,
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

		hubOrganizationName, hubRepositoryName, err := parseRemoteURL(gitRemoteURL)
		if err != nil {
			log.Fatal(err)
		}

		// Get repository info from Github API
		ctx := context.Background()
		client := github.NewClient(nil)
		hubRepository, _, err := client.Repositories.Get(ctx, hubOrganizationName, hubRepositoryName)
		if err != nil {
			log.Fatal(err)
		}

		// Print info
		fmt.Printf("GitHub Repository ID: %d\n", hubRepository.ID)
		fmt.Printf("Github Respository Full Name: %s\n", *hubRepository.FullName)
		fmt.Printf("Github Respository URL: %s\n", *hubRepository.HTMLURL)
	},
}

func parseRemoteURL(url string) (string, string, error) {
	parts := strings.Split(url, ":")
	parts = strings.Split(parts[1], ".")
	parts = strings.Split(parts[0], "/")
	organization := parts[0]
	repository := parts[1]
	return organization, repository, nil
}
