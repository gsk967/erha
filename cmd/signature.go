package cmd

import (
	"fmt"

	"github.com/gsk967/erha/fsoper"
	"github.com/gsk967/erha/signature"
	"github.com/spf13/cobra"
)

// GenSigCmd returns cmd for generating signature for given file.
func GenSigCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:     "signature",
		Short:   "Generate signature of the file.",
		Args:    cobra.ExactArgs(2),
		Long:    `It will generate signature for given file using rooling hash algorithm.`,
		Example: "erha signature filename.txt signature.txt --chunk_size 4",
		RunE: func(cmd *cobra.Command, args []string) error {
			chunkSize, err := cmd.Flags().GetInt64(FlagChunkSize)
			if err != nil {
				return err
			}
			if !fsoper.FileExists(args[0]) {
				return fmt.Errorf("file is not exists %s , Please check this", args[0])
			}

			return signature.Generate(args[0], args[1], chunkSize)
		},
	}
	return &cmd
}
