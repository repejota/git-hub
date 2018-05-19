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

	"github.com/google/go-github/github"
)

// ListIssuesByRepo ...
func ListIssuesByRepo(repoFullName string) ([]*github.Issue, error) {
	organization, repository := ParseFullName(repoFullName)
	ctx := context.Background()
	client := github.NewClient(nil)
	options := &github.IssueListByRepoOptions{}
	issues, _, err := client.Issues.ListByRepo(ctx, organization, repository, options)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
