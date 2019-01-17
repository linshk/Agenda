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

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove -t [title] -p [participators]",
	Short: "remove participators from the meeting",
	Long: `remove participators from the meeting specified by the title, the meeting will be canceled if there is no participator after removing`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("remove --title=%s --participators=%s", title, participators)
	
		if err := entity.RemoveParticipators(title, participators); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Println("remove participators successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	removeCmd.Flags().StringP("title","t","","title of meeting")
	removeCmd.MarkFlagRequired("title")
	
	removeCmd.Flags().StringP("participators","p","","participators of the meeting")
	removeCmd.MarkFlagRequired("participators")
}
