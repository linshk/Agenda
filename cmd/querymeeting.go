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
	"github.com/modood/table"
)

// querymeetingCmd represents the querymeeting command
var querymeetingCmd = &cobra.Command{
	Use:   "querymeeting",
	Short: "query meetings by time interval",
	Long: `query meetings that current user participates or sponsors by time interval`,
	Run: func(cmd *cobra.Command, args []string) {
		st, _ := cmd.Flags().GetString("starttime")
		et, _ := cmd.Flags().GetString("endtime")

		Log.SetPrefix("[Cmd]   ")
		Log.Printf("querymeeting  --starttime=%v --endtime=%v", st, et)
	
		if err, meetings := entity.QueryMeetingsByTime(st, et); err != nil {
			Log.SetPrefix("[Error] ")
			Log.Println(err)
		} else {
			Log.SetPrefix("[OK]    ")
			Log.Print("\n" + table.Table(meetings))
		}
	},
}

func init() {
	rootCmd.AddCommand(querymeetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// querymeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// querymeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")	
	querymeetingCmd.Flags().StringP("starttime","s","","start time")
	querymeetingCmd.MarkFlagRequired("starttime")

	querymeetingCmd.Flags().StringP("endtime","e","","end time")
	querymeetingCmd.MarkFlagRequired("endtime")
}
