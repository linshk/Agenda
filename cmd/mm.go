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
	"fmt"

	"github.com/spf13/cobra"
)

// mmCmd represents the mm command
var mmCmd = &cobra.Command{
	Use:   "mm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {	
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetString("participator")
		remove, _ := cmd.Flags().GetBool("remove")
		fmt.Printf("mm called with args: title=%s participators=%s remove=%v\n",title,participators,remove)
	},
}

func init() {
	rootCmd.AddCommand(mmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mmCmd.Flags().StringP("title","t","","title of meeting")
	mmCmd.MarkFlagRequired("title")
	mmCmd.Flags().StringP("participator","p","","participators to add/romove")
	mmCmd.MarkFlagRequired("participator")
	mmCmd.Flags().BoolP("remove","r",false,"remove participators")
}
