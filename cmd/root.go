// Copyright © 2019 John F. Hogarty <hogihung@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var filterBy string
var targetsFile string
var logFile string
var executablePath string

// NOTE: may need to refactor to pull code/checks for the targets JSON file into
//       the files files test.go and list.go.  The adhoc subcommand does not need
//       the targets JSON file.  Only the test and list subcommands need it.

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra_kai",
	Short: "Command line utility for running REST GET requests against target endpoints",
	Long: `Cobra Kai is a command line utility, which takes a JSON formatted configuration
file and performs REST GET requests against the defined target endpoints. 

Using this app, with JSON formatted config file, one can run n number of GET requests to the
defined target endpoints and display the response to the terminal (default) and/or write the
responses to a file.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	executablePath = filepath.Dir(ex)

	cobra.OnInitialize(initSetup)

	rootCmd.PersistentFlags().StringVar(&logFile, "log", "", "log file (default is ./rest_easy.log)")
}

func initSetup() {

	// Checks for the targets JSON file (required unless adhoc subcommand is used.)
	if targetsFile == "" {
		targetsFile = executablePath + "/targets.json"
	}

	if fileExists(targetsFile) {
		targetsFileFH, err := ioutil.ReadFile(targetsFile)
		if err != nil {
			log.Fatal(err)
		}

		// Next we would make sure the readable file is valid JSON format
		if !json.Valid(targetsFileFH) {
			log.Fatalf("File %v is not a valid JSON file.", targetsFile)
		}
	} else {
		log.Fatalf("File does not exist: %v, exiting.", targetsFile)
	}

	// Checks for the log file
	if logFile == "" {
		logFile = executablePath + "/rest_easy.log"
	}

	if fileExists(logFile) {
		logFileFH, err := ioutil.ReadFile(logFile)
		if err != nil {
			fmt.Println("WARNING:  Unable to read log file.  Operation will continue, but no log will be written.")
		}
		err = ioutil.WriteFile(logFile, logFileFH, 0644)
		if err != nil {
			fmt.Println("WARNING: File is NOT writable.  Operation will continue but no log will be written.")
		}
	} else {
		fmt.Println("WARNING:  Log file does not exist. Operation will continue, but no log will be written.")
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// URLTarget - Struct to map data from targetsFile
type URLTarget struct {
	Targets []struct {
		URL   string `json:"url"`
		Auth  string `json:"auth,omitempty"`
		User  string `json:"user,omitempty"`
		Pass  string `json:"pass,omitempty"`
		Label string `json:"label,omitempty"`
		Group string `json:"group,omitempty"`
		Token string `json:"token,omitempty"`
	} `json:"targets"`
}
