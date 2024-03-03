package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/gsk967/erha/delta"
	"github.com/gsk967/erha/fsoper"
)

// DeltaCmd returns cmd for generate delta of two files
func DeltaCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delta",
		Short: "Delta for the files.",
		Args:  cobra.ExactArgs(3),
		Long: `This command displays the copied, updated, and deleted contents of the original file relative to the new file. 
		It utilizes the rolling hash file diffing algorithm.`,
		Example: "erha delta signature-file.txt new-file.txt delta-file.txt",
		RunE: func(cmd *cobra.Command, args []string) error {
			if !fsoper.FileExists(args[0]) {
				return fmt.Errorf("signature file is not exists %s , Please check this", args[0])
			}
			if !fsoper.FileExists(args[1]) {
				return fmt.Errorf("new file is not exists %s , Please check this", args[0])
			}

			return delta.GenerateDelta(args[0], args[1], args[2])
		},
	}
	return &cmd
}
