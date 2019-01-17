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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add -t [title] -p [participators]",
	Short: "add participators to the meeting",
	Long: `add participators to the meeting that you sponsored`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("add --title=%s --participators=%s", title, participators)

		if err := entity.AddParticipators(title, participators); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Println("add participators successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().StringP("title","t","","title of meeting")
	addCmd.MarkFlagRequired("title")
	
	addCmd.Flags().StringP("participators","p","","participators of the meeting")
	addCmd.MarkFlagRequired("participators")
}
