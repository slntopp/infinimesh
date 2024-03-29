//--------------------------------------------------------------------------
// Copyright 2018 infinimesh
// www.infinimesh.io
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.
//--------------------------------------------------------------------------

package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/slntopp/infinimesh/pkg/node/nodepb"
)

func init() {
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountListCmd)
	rootCmd.AddCommand(accountCmd)
}

var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return connectGRPC()
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return disconnectGRPC()
	},
}

var accountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create user account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := accountClient.CreateUserAccount(ctx, &nodepb.CreateUserAccountRequest{
			Account: &nodepb.Account{
				Name:     args[0],
				Enabled:  true,
				Password: args[1],
			},
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create user: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Created user %v with id %v.\n", args[0], response.GetUid())
	},
}

var accountListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List user accounts",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		defer w.Flush()

		response, err := accountClient.ListAccounts(ctx, &nodepb.ListAccountsRequest{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list accounts: %v\n", err)
			os.Exit(1)
		}

		if !noHeaderFlag {
			fmt.Fprintf(w, "NAME\tID\t\n")
		}

		for _, account := range response.GetAccounts() {
			fmt.Fprintf(w, "%v\t%v\n", account.GetName(), account.GetUid())
		}

	},
}
