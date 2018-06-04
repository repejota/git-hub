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

package automation

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/repejota/git-hub"
)

// GetCurrentBranch ...
func GetCurrentBranch(repository *ghub.Repository) (string, error) {
	out, err := exec.Command("git", "symbolic-ref", "--short", "HEAD").Output()
	if err != nil {
		return "", err
	}
	sout := strings.Trim(string(out), "\n")
	return sout, nil
}

// PullMasterBranch ...
func PullMasterBranch(repository *ghub.Repository) (string, error) {
	out, err := exec.Command("git", "pull", "origin", "master").Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// CreateLocalGitBranch ...
func CreateLocalGitBranch(name string) (string, error) {
	out, err := exec.Command("git", "checkout", "-b", name).Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// PushLocalBranchToOrigin ...
func PushLocalBranchToOrigin(name string) (string, error) {
	out, err := exec.Command("git", "push", "--set-upstream", "origin", name).Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// BumpNextVersion ...
func BumpNextVersion(nextversion *ghub.SemVer) (string, error) {
	finalOut := ""

	// update VERSION file contents
	data := []byte(nextversion.String())
	sdata := strings.Trim(string(data), "\n")
	err := ioutil.WriteFile("VERSION", []byte(sdata), 0644)
	if err != nil {
		return "", nil
	}

	// commit VERSION changes
	out, err := exec.Command("git", "add", "VERSION").Output()
	if err != nil {
		return "", err
	}
	finalOut = fmt.Sprintf("%s%s", finalOut, string(out))

	commitMsg := fmt.Sprintf("Bump %s", nextversion)
	out, err = exec.Command("git", "commit", "VERSION", "-m", commitMsg).Output()
	if err != nil {
		return "", err
	}
	finalOut = fmt.Sprintf("%s%s", finalOut, string(out))

	// push VERSION bump commit
	sout, err := GitPush()
	if err != nil {
		return "", err
	}
	finalOut = fmt.Sprintf("%s%s", finalOut, sout)

	return finalOut, nil
}

// GitPush ...
func GitPush() (string, error) {
	out, err := exec.Command("git", "push").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// GoGitBranch ...
func GoGitBranch(name string) (string, error) {
	out, err := exec.Command("git", "checkout", name).Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// PullAndRebase ...
func PullAndRebase() (string, error) {
	out, err := exec.Command("git", "pull", "--rebase", "--prune").Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// MergeBranch ...
func MergeBranch(branchName string) (string, error) {
	out, err := exec.Command("git", "merge", "--no-ff", "--no-edit", branchName).Output()
	if err != nil {
		return "", err
	}
	sout := string(out)
	return sout, nil
}

// CreateGitTag ...
func CreateGitTag(tagName string) (string, error) {
	msgTag := fmt.Sprintf("Release %s", tagName)
	out, err := exec.Command("git", "tag", "-a", tagName, "-m", msgTag).Output()
	if err != nil {
		return "", nil
	}
	sout := string(out)
	return sout, nil
}

// GitPushTags ...
func GitPushTags() (string, error) {
	out, err := exec.Command("git", "push", "--tags").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// DeleteRemoteBranch ...
func DeleteRemoteBranch(branchName string) (string, error) {
	out, err := exec.Command("git", "push", "origin", "-d", branchName).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// DeleteLocalBranch ...
func DeleteLocalBranch(branchName string) (string, error) {
	out, err := exec.Command("git", "branch", "-d", branchName).Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
