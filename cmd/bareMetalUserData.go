// Copyright © 2019 The Vultr-cli Authors
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
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vultr/vultr-cli/cmd/printer"
)

// BareMetalUserData represents the baremetal userdata commands
func BareMetalUserData() *cobra.Command {
	bareMetalUserDataCmd := &cobra.Command{
		Use:     "user-data",
		Short:   "user-data is used to access bare metal server user-data commands",
		Aliases: []string{"u"},
	}

	bareMetalUserDataCmd.AddCommand(bareMetalGetUserData)

	return bareMetalUserDataCmd
}

var bareMetalGetUserData = &cobra.Command{
	Use:     "get <bareMetalID>",
	Short:   "Get the user-data of a bare metal server.",
	Aliases: []string{"g"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a bareMetalID")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		u, err := client.BareMetalServer.GetUserData(context.Background(), args[0])
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		printer.UserData(u)
	},
}
