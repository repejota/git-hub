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
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// ReleaseCmd represents the release command
var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Manage releases",
	Long:  `Manage repository and Github releases`,
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
		log.Printf("GitHub Token: %s\n", gitHubToken)

		cmd.Usage()
		os.Exit(0)
	},
}
