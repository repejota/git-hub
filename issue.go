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
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// ListIssuesByRepo ...
func ListIssuesByRepo(repoFullName string) ([]*github.Issue, error) {
	organization, repository := ParseRepositoryFullName(repoFullName)
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	options := &github.IssueListByRepoOptions{}
	issues, _, err := client.Issues.ListByRepo(ctx, organization, repository, options)
	if err != nil {
		return nil, err
	}
	return issues, nil
}

// GetIssue ...
func GetIssue(organization string, repository string, issueID int) (*github.Issue, error) {
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	issue, _, err := client.Issues.Get(ctx, organization, repository, issueID)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

// AssignUserToIssue ...
func AssignUserToIssue(organization string, repository string, user *github.User, issue *github.Issue) error {
	ctx := context.Background()
	githubToken := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: githubToken,
		},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	users := []string{*user.Login}
	issue, _, err := client.Issues.AddAssignees(ctx, organization, repository, *issue.Number, users)
	if err != nil {
		return err
	}
	return nil
}

// SlugifyIssue ...
func SlugifyIssue(issue *github.Issue) string {
	var re = regexp.MustCompile("[^a-z0-9]+")
	slugifyTitle := strings.Trim(re.ReplaceAllString(strings.ToLower(issue.GetTitle()), "-"), "-")
	slugIssue := fmt.Sprintf("%d-%s", issue.GetNumber(), slugifyTitle)
	return slugIssue
}

// SlugifyRepository ...
func SlugifyRepository(repositoryFullName string) string {
	var re = regexp.MustCompile("[^a-z0-9]+")
	slugifyRepo := strings.Trim(re.ReplaceAllString(strings.ToLower(repositoryFullName), "-"), "-")
	return slugifyRepo
}
