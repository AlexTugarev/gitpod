// Copyright (c) 2020 TypeFox GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package cmd

import (
	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "A number of shortcuts to the inner workings of wsman to debug its functionality",
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
