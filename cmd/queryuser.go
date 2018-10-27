// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/linshk/Agenda/entity"
	"github.com/spf13/cobra"
)

// queryuserCmd represents the register command
var queryuserCmd = &cobra.Command{
	Use:   "queryuser",
	Short: "view all registered users",
	Long:  "view all registered users with username, email and phone",
	Run: func(cmd *cobra.Command, args []string) {

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("queryuser")

		if userList, err := entity.QueryUser(); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Print(userList)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryuserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
