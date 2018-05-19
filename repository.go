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
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	git "gopkg.in/src-d/go-git.v4"
)

// Repository ...
type Repository struct {
	GitRepository    *git.Repository
	GitHubRepository *github.Repository
}

// Git ...
func (r *Repository) Git(path string) error {
	gitRepository, err := git.PlainOpen(path)
	if err != nil {
		return err
	}
	r.GitRepository = gitRepository
	return nil
}

// GetRemoteGithubRepository ...
func (r *Repository) GetRemoteGithubRepository(remoteName string) error {
	// Get remote
	remote, err := r.GitRepository.Remote(remoteName)
	if err != nil {
		return err
	}
	remoteURL := remote.Config().URLs[0]
	host, organization, repository, err := ParseGithubURL(remoteURL)
	if err != nil {
		return nil
	}
	if host != "github.com" {
		return errors.New("Remote host is not github.com")
	}
	// Get repository info from Github API
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	githubRepository, _, err := client.Repositories.Get(ctx, organization, repository)
	if err != nil {
		log.Fatal(err)
	}
	r.GitHubRepository = githubRepository
	return nil
}

// ParseGithubURL ...
func ParseGithubURL(url string) (string, string, string, error) {
	// git@github.com:repejota/git-hub.git
	parts := strings.Split(url, ":")
	host := strings.Split(parts[0], "@")[1]
	fullName := strings.Split(parts[1], ".")[0]
	organization, repository := ParseFullName(fullName)
	return host, organization, repository, nil
}

// ParseFullName ...
func ParseFullName(fullName string) (string, string) {
	parts := strings.Split(fullName, "/")
	organization := parts[0]
	repository := parts[1]
	return organization, repository
}
