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

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "create -t [title] -p [participators] -s [start time] -e [end time]",
	Short: "create a meeting",
	Long: `create a meeting, with a unique title, a list of participators (a list of registered usernames separated by comma), and with start time and end time (format:2006-01-02T15:04:05).
	Meetings with conflicts are not allowed`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participators")
		st, _ := cmd.Flags().GetString("starttime")
		et, _ := cmd.Flags().GetString("endtime")

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("create --title=%s --participators=%s --starttime=%v --endtime=%v", title, participators, st, et)
	
		if err := entity.CreateMeeting(title, participators, st, et); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Println("create meeting successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmCmd.Flags().StringP("title","t","","title of meeting")
	cmCmd.MarkFlagRequired("title")
	
	cmCmd.Flags().StringP("participators","p","","participators of the meeting")
	cmCmd.MarkFlagRequired("participators")
	
	cmCmd.Flags().StringP("starttime","s","","start time")
	cmCmd.MarkFlagRequired("starttime")

	cmCmd.Flags().StringP("endtime","e","","end time")
	cmCmd.MarkFlagRequired("endtime")
}
