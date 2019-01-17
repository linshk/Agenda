// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

// quitCmd represents the quit command
var quitCmd = &cobra.Command{
	Use:   "quit -t [title]",
	Short: "quit a meeting participated by current user",
	Long: `quit a meeting participated by current user, the meeting will be canceled if there is no participator after quiting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("quit --title=%s", title)
	
		if err := entity.QuitMeeting(title); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Println("quit meeting successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(quitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	quitCmd.Flags().StringP("title","t","","title of meeting")
	quitCmd.MarkFlagRequired("title")
}
