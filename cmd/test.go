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
	"io/ioutil"
	"net/http"

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
		Logr.Info("-- Test sub-command called")

		all := cmd.Flag("all").Value.String()
		if all == "true" {
			filterBy = "all"
			testFilter := Filter(filterBy, all)
			executeTest(testFilter)
		}

		group := cmd.Flag("group").Value.String()
		if group != "" {
			filterBy = "group"
			testFilter := Filter(filterBy, group)
			executeTest(testFilter)
		}

		selection := cmd.Flag("selection").Value.String()
		if selection != "" {
			filterBy = "selection"
			testFilter := Filter(filterBy, selection)
			executeTest(testFilter)
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

func executeTest(targets URLTargets) {
	for i := 0; i < len(targets.Target); i++ {
		Logr.Info("Preparing to execute a request to url: ", targets.Target[i].URL)

		if isNoneAuth(targets.Target[i].Auth) {
			executeNoneAuthGet(targets.Target[i].URL)
		}

		if isBasicAuth(targets.Target[i].Auth) {
			executeBasicAuthGet(targets.Target[i].URL, targets.Target[i].User, targets.Target[i].Pass)
		}

		if isTokenAuth(targets.Target[i].Auth) {
			executeTokenAuthGet(targets.Target[i].URL, targets.Target[i].Token)
		}
		//Logr.Warn("Failed to determine auth type for URL ", targets.Target[i].URL, ", moving on")
	}
}

func executeNoneAuthGet(url string) {
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	httpResponse, err := client.Do(req)

	if err != nil {
		Logr.Warn(err)
		return
	}
	defer httpResponse.Body.Close()

	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		Logr.Warn(err)
		return
	}

	result := HTTPResponse{httpResponse.Status, httpBody}
	Logr.WithFields(Logr.Fields{
		"status": result.status,
	}).Info("Request completed for: ", url)
}

func executeBasicAuthGet(url string, user string, password string) {
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, password)

	client := &http.Client{}
	httpResponse, err := client.Do(req)

	if err != nil {
		Logr.Warn(err)
		return
	}
	defer httpResponse.Body.Close()

	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		Logr.Warn(err)
		return
	}

	result := HTTPResponse{httpResponse.Status, httpBody}
	Logr.WithFields(Logr.Fields{
		"status": result.status,
	}).Info("Request completed for: ", url)

}

func executeTokenAuthGet(url string, token string) {
	var bearer = "Bearer " + token

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	httpResponse, err := client.Do(req)

	if err != nil {
		Logr.Warn(err)
		return
	}
	defer httpResponse.Body.Close()

	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		Logr.Warn(err)
		return
	}

	result := HTTPResponse{httpResponse.Status, httpBody}
	Logr.WithFields(Logr.Fields{
		"status": result.status,
	}).Info("Request completed for: ", url)

}

func isBasicAuth(auth string) bool {
	if auth == "basic" {
		return true
	}
	return false
}

func isTokenAuth(auth string) bool {
	if auth == "token" {
		return true
	}
	return false
}

func isNoneAuth(auth string) bool {
	if auth == "none" {
		return true
	}
	if isBasicAuth(auth) {
		return false
	}
	if isTokenAuth(auth) {
		return false
	}
	return false
}

// NOTE: I'm not happy with this.  Need to find a better way so that we can
//       create log entry when auth type is not determinable
func isUnknownAuth(auth string) bool {
	if auth == "none" {
		return false
	}
	if isBasicAuth(auth) {
		return false
	}
	if isTokenAuth(auth) {
		return false
	}
	return true
}

// HTTPResponse is a struct for handling the responses we will be getting from
// our GET requests.
type HTTPResponse struct {
	status string
	body   []byte
}
