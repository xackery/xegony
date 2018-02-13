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
	"time"

	"github.com/spf13/cobra"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// zoneImageCmd represents the servzone imageer command
var zoneImageCmd = &cobra.Command{
	Use:   "zoneimage",
	Short: "Run bot for zone image",
	Long:  `Run zone image task`,
	Run:   runZoneImage,
}

func init() {
	rootCmd.AddCommand(zoneImageCmd)

}

func runZoneImage(cmd *cobra.Command, args []string) {

	err := startZoneImage(cmd, args)
	if err != nil {
		fmt.Println("fail during zoneimage:", err)
	}
}

func startZoneImage(cmd *cobra.Command, args []string) (err error) {
	start := time.Now()
	_, _, _, _, _, err = initializeSystem()
	if err != nil {
		return
	}
	log.Printf("Starting zone image...")
	user := &model.User{
		PrimaryAccount: &model.Account{
			Status: 200,
		},
	}

	bot := &model.Bot{
		ID: 1,
	}
	err = cases.GetZoneImageBot(bot, user)
	if err != nil {
		return
	}

	bot.Status = 1
	err = cases.EditZoneImageBot(bot, user)
	if err != nil {
		return
	}

	log.Printf("Completed in %s\n", time.Since(start))
	return
}
