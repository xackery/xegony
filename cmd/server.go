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
	alog "log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/bot"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/storage"
	"github.com/xackery/xegony/storage/mariadb"
	"github.com/xackery/xegony/web"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Start the server based on provided parameters. By default, it will run API, BOT, and WEB servers.`,
	Run:   runServer,
}

var connection string
var dbtype string
var log *alog.Logger
var logErr *alog.Logger

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringVarP(&dbtype, "dbtype", "", "mysql", "type of database to connect to")
}

func runServer(cmd *cobra.Command, args []string) {

	err := startServer(cmd, args)
	if err != nil {
		fmt.Println("fail during startserver:", err)
	}
}

func startServer(cmd *cobra.Command, args []string) (err error) {
	start := time.Now()
	var sw storage.Writer
	var sr storage.Reader
	var si storage.Initializer
	w := os.Stdout
	wErr := os.Stderr

	log = alog.New(w, "Main: ", 0)
	logErr = alog.New(wErr, "MainErr: ", 0)

	log.Printf("Loading data to memory...")
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	var totalMemoryInUse uint64
	totalMemoryInUse = m.Alloc

	//We start with the config, since other endpoints utilize this information.
	err = cases.LoadConfigFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load config to memory")
		return
	}

	//parse arguments now that we have config info
	if dbtype == "mysql" {
		var db *mariadb.Storage
		db, err = mariadb.New(cases.GetConfigForMySQL(), nil, nil)
		if err != nil {
			log.Fatal("Failed to create mariadb:", err.Error())
		}
		sw = db
		sr = db
		si = db
	} else {
		log.Fatal("unsupported db type:", dbtype)
	}

	err = si.VerifyTables()
	if err != nil {
		log.Fatal("Failed to verify tables: ", err.Error())
	}

	err = cases.InitializeAll(sr, sw, si)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize all")
		return
	}

	err = cases.LoadClassFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = cases.LoadDeityFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = cases.LoadRaceFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = cases.LoadRuleFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load rule to memory")
		return
	}

	err = cases.LoadRuleEntryFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load ruleEntry to memory")
		return
	}

	err = cases.LoadSpellAnimationFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellAnimation to memory")
		return
	}

	err = cases.LoadSpellAnimationTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellAnimationType to memory")
		return
	}

	err = cases.LoadSpellDurationFormulaFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellDurationFormula to memory")
		return
	}

	err = cases.LoadSpellEffectFormulaFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellEffectFormula to memory")
		return
	}

	err = cases.LoadSpellEffectTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellEffectType to memory")
		return
	}

	err = cases.LoadSpellTravelTypeFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load spellTravelType to memory")
		return
	}

	err = cases.LoadVariableFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load variable to memory")
		return
	}

	err = cases.LoadZoneFromDBToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zone to memory")
		return
	}

	err = cases.LoadZoneExpansionFromFileToMemory()
	if err != nil {
		err = errors.Wrap(err, "failed to load zoneExpansion to memory")
		return
	}

	fmt.Printf("\n")
	runtime.ReadMemStats(m)
	if m.Alloc > totalMemoryInUse {
		totalMemoryInUse = m.Alloc - totalMemoryInUse
	} else {
		totalMemoryInUse = 0
	}

	router := mux.NewRouter().StrictSlash(true)

	if err = api.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize api:", err.Error())
	}
	api.ApplyRoutes(router)

	if err = bot.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize bot:", err.Error())
	}
	bot.ApplyRoutes(router)

	if err = web.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize web:", err.Error())
	}
	web.ApplyRoutes(router)

	log.Println("Listening on", cases.GetConfigForHTTP())

	log.Printf("Started in %s. Using %s for data in memory, %s total\n", time.Since(start), humanize.Bytes(totalMemoryInUse), humanize.Bytes(m.TotalAlloc))
	err = http.ListenAndServe(cases.GetConfigForHTTP(), router)
	return
}
