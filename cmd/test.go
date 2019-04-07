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
	"fmt"

	Logr "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Use the 'test' sub-command to run GET requests to defined endpoints",
	Long: `Use the 'test' sub-command, with required and/or optional flags and
arguments to perform REST GET requests against the endpoint targets defined
in the JSON file.

With the 'test' sub-command one can run tests against all (default)
endpoints, against a group of endpoints, or against a selection of
defined endpoints.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("---- test called ----")
		Logr.Info("-- Test sub-command called")

		all := cmd.Flag("all").Value.String()
		if all == "true" {
			filterBy = "all"
			fmt.Println("Filter set to: ", filterBy)
			return
		}

		group := cmd.Flag("group").Value.String()
		if group != "" {
			filterBy = "group"
			fmt.Println("Filter set to: ", filterBy)
			fmt.Println("With a group value of: ", group)
			return
		}

		selection := cmd.Flag("selection").Value.String()
		if selection != "" {
			filterBy = "selection"
			fmt.Println("Filter set to: ", filterBy)
			fmt.Println("With a selection of: ", selection)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	testCmd.Flags().BoolP("all", "a", false, "Use to test all items in JSON file")
	testCmd.Flags().String("group", "", "Use for a targeted group")
	testCmd.Flags().String("selection", "", "Use for a selection of items - 'labels'")
	testCmd.PersistentFlags().StringVar(&targetsFile, "targets", "", "JSON formatted targets file (default is ./targets.json)")
}
