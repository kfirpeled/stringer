package stringer

import (
	"fmt"

	"github.com/kfirpeled/stringer/pkg/stringer"
	"github.com/spf13/cobra"
)

func NewReverseCmd() *cobra.Command {
	var reverseCmd = &cobra.Command{
		Use:     "reverse",
		Aliases: []string{"rev"},
		Short:   "Reverses a string",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			res := stringer.Reverse(args[0])
			fmt.Fprintln(cmd.OutOrStdout(), res)
			return nil
		},
	}
	return reverseCmd
}

func init() {
	rootCmd.AddCommand(NewReverseCmd())
}
