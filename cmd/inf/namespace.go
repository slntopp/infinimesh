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
	listNamespacesCmd.Flags().BoolVar(&noHeaderFlag, "no-headers", false, "Hide table headers")
	namespaceCmd.AddCommand(describeNamespace)
	namespaceCmd.AddCommand(listNamespacesCmd)
	namespaceCmd.AddCommand(createNamespaceCmd)
	rootCmd.AddCommand(namespaceCmd)

}

var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns", "namespaces"},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return connectGRPC()
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return disconnectGRPC()
	},
}

var createNamespaceCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_, err := namespaceClient.CreateNamespace(ctx, &nodepb.CreateNamespaceRequest{Name: args[0]})
		if err != nil {
			fmt.Println("grpc: failed to create namespace", err)
			os.Exit(1)
		}
		fmt.Printf("Created namespace %v.\n", args[0])
	},
}

var listNamespacesCmd = &cobra.Command{
	Use:     "list",
	Short:   "List namespaces",
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)
		defer w.Flush()

		response, err := namespaceClient.ListNamespaces(ctx, &nodepb.ListNamespacesRequest{})
		if err != nil {
			fmt.Println("grpc: failed to fetch data", err)
			os.Exit(1)
		}
		_ = response
		if !noHeaderFlag {
			fmt.Fprintf(w, "NAME\tID\t\n")
		}

		for _, ns := range response.Namespaces {
			fmt.Fprintf(w, "%v\t%v\n", ns.GetName(), ns.GetId())
		}
	},
}

var describeNamespace = &cobra.Command{
	Use:     "describe",
	Short:   "Describe namespace",
	Aliases: []string{"desc"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("desc")
		w := tabwriter.NewWriter(os.Stdout, tabwriterMinWidth, tabwriterWidth, tabwriterPadding, tabwriterPadChar, tabwriterFlags)

		response, err := namespaceClient.GetNamespace(ctx, &nodepb.GetNamespaceRequest{Namespace: args[0]})
		if err != nil {
			fmt.Println("grpc: failed to fetch data", err)
			os.Exit(1)
		}
		fmt.Println(response)
		// _ = response
		// if !noHeaderFlag {
		// 	fmt.Fprintf(w, "NAME\tID\t\n")
		// }

		// for _, ns := range response.Namespaces {
		// 	fmt.Fprintf(w, "%v\t%v\n", ns.GetName(), ns.GetId())
		// }

		defer w.Flush()
	},
}
