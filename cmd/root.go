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
	"os"
	"runtime"
	"time"

	"github.com/dustin/go-humanize"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/storage"
	"github.com/xackery/xegony/storage/mariadb"
)

var cfgFile string
var eqemuconfig string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "xegony",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Println("Exited successfully")
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.xegony.yaml)")
	//rootCmd.PersistentFlags().StringP("eqemuconfig", "", "YOUR NAME", "Author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&eqemuconfig, "eqemuconfig", "", "./eqemu_config.xml", "where to find eqemu_config.xml")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".xegony" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".xegony")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initializeSystem() (sw storage.Writer, sr storage.Reader, si storage.Initializer, w *os.File, wErr *os.File, err error) {
	start := time.Now()
	w = os.Stdout
	wErr = os.Stderr

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

	err = cases.InitializeAllDatabaseStorage(sr, sw, si)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize all")
		return
	}

	err = cases.InitializeAllMemoryStorage()
	if err != nil {
		err = errors.Wrap(err, "failed to initialize memory story")
		return
	}

	err = cases.InitializeAllWorkers()
	if err != nil {
		err = errors.Wrap(err, "failed to initialize workers")
		return
	}

	fmt.Printf("\n")
	runtime.ReadMemStats(m)
	if m.Alloc > totalMemoryInUse {
		totalMemoryInUse = m.Alloc - totalMemoryInUse
	} else {
		totalMemoryInUse = 0
	}

	log.Printf("Loaded in %s. Using %s for data in memory, %s total\n", time.Since(start), humanize.Bytes(totalMemoryInUse), humanize.Bytes(m.TotalAlloc))
	return
}
