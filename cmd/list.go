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
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use the 'list' sub-command to list the defined endpoints to be tested",
	Long: `Use the 'list' sub-command, with required and/or optional flags and
arguments to list the endpoints to which REST GET requests would be sent against, as defined
in the JSON file.

With the 'list' sub-command one can see the all (default) the endpoints to test,
a group of endpoints to test, or a selection of endpoints to test.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---- list called ----")

		all := cmd.Flag("all").Value.String()
		if all == "true" {
			filterBy = "all"
			fmt.Println("Filter set to: ", filterBy)
			// This should print all of our records extracted from targets.json
			// See the print example below.
			display(all)
			return
		}

		group := cmd.Flag("group").Value.String()
		if group != "" {
			filterBy = "group"
			fmt.Println("Filter set to: ", filterBy)
			fmt.Println("With a group value of: ", group)
			// Here we should only print records extracted from targets.json
			// that contain a value for group that matches what was supplied
			// via the command line.
			return
		}

		selection := cmd.Flag("selection").Value.String()
		if selection != "" {
			filterBy = "selection"
			fmt.Println("Filter set to: ", filterBy)
			fmt.Println("With a selection of: ", selection)
			// Here we should only print records extracted from targets.json
			// that contain a value(s) for label that matches what was supplied
			// for selection via the command line.  Can be one or more values.
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("all", "a", false, "Use to test all items in JSON file")
	listCmd.Flags().String("group", "", "Use for a targeted group")
	listCmd.Flags().String("selection", "", "Use for a selection of items - 'labels'")
	listCmd.PersistentFlags().StringVar(&targetsFile, "targets", "", "JSON formatted targets file (default is ./targets.json)")
}

func display(filter string) {
	if filter == "true" {
		fmt.Println("Preparing to parse JSON for file: ", targetsFile)
		jsonFile, err := os.Open(targetsFile)
		if err != nil {
			fmt.Println(err) // change this to log.Fatal or panic
		}
		defer jsonFile.Close()

		jsonValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err) // change this to log.Fatal or panic
		}

		fmt.Println("-----------------------------------")
		targets := URLTarget{}
		_ = json.Unmarshal([]byte(jsonValue), &targets)

		for i := 0; i < len(targets.Targets); i++ {
			fmt.Println("URL: " + targets.Targets[i].URL)
			fmt.Println("Auth: " + targets.Targets[i].Auth)
			fmt.Println("User: " + targets.Targets[i].User)
			fmt.Println("Pass: " + targets.Targets[i].Pass)
			fmt.Println("Label: " + targets.Targets[i].Label)
			fmt.Println("Group: " + targets.Targets[i].Group)
			fmt.Println("Token: " + targets.Targets[i].Token)
			fmt.Println("-------------------------------------------")

			// NOTE: above we are just getting the feel of things and how to
			//       dissect our JSON data.  Eventually we will remove the
			//       Printlns and replace with the appropriate function.
			//
			// https://gobyexample.com/switch
		}
	}
}
