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

// adhocCmd represents the adhoc command
var adhocCmd = &cobra.Command{
	Use:   "adhoc",
	Short: "Use the 'adhoc' sub-command to run GET requests against single endpoint",
	Long: `With the 'adhoc' sub-command, and supplied required and optional flags
and arguments one can execute a REST GET request againt the target endpoint.

The 'adhoc' command takes a required flag and argument for endpoint along with
optional flags/arguments for authentication.`,
	Run: func(cmd *cobra.Command, args []string) {
		Logr.Info("-- Adhoc sub-command called")

		endpoint := cmd.Flag("endpoint").Value.String()
		fmt.Println("Endpoint: ", endpoint)

		auth := cmd.Flag("auth").Value.String()
		fmt.Println("Auth: ", auth)

		user := cmd.Flag("user").Value.String()
		fmt.Println("User: ", user)

		pass := cmd.Flag("pass").Value.String()
		fmt.Println("Password: ", pass)

		token := cmd.Flag("token").Value.String()
		fmt.Println("Token: ", token)
	},
}

func init() {
	rootCmd.AddCommand(adhocCmd)

	adhocCmd.Flags().String("endpoint", "", "Enter target endpoint to be tested (required)")
	adhocCmd.Flags().String("auth", "none", "Use either basic, token or none (default is none)")
	adhocCmd.Flags().String("user", "", "Use with auth type 'basic' along with pass")
	adhocCmd.Flags().String("pass", "", "Use with auth type 'basic' along with user")
	adhocCmd.Flags().String("token", "", "Use with auth type 'token'")
}
