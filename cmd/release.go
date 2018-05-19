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

// ReleaseCmd represents the release command
var ReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Get information releases",
	Long:  `Get information about the repository releases`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
		os.Exit(0)
	},
}

// ReleaseListCmd represents the release list command
var ReleaseListCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases",
	Long:  `List releases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub release list")
	},
}

// ReleaseStartCmd represents the release start command
var ReleaseStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a release",
	Long:  `Start working on an release`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub release start")
	},
}

// ReleaseFinishCmd represents the release finish command
var ReleaseFinishCmd = &cobra.Command{
	Use:   "finish",
	Short: "Finish an release",
	Long:  `Finish working on an release`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git hub release finish")
	},
}