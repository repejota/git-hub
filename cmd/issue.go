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
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub list issue")
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
		fmt.Println("git hub list finish")
	},
}
