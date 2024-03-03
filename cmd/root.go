/*
Copyright © 2024 Sai Kumar <gskumar967@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/gsk967/erha/fsoper"
	"github.com/gsk967/erha/types"
	"github.com/spf13/cobra"
)

// Note: It will not write the signature and delta of files into output files.
// For this, please use subcommands of erha.
var rootCmd = &cobra.Command{
	Use:     "erha",
	Short:   "files diff based on rolling hash algorithm.",
	Args:    cobra.ExactArgs(2),
	Long:    `It is a rolling hash algorithm for file differencing.`,
	Example: "erha original-file.txt new-file.txt",
	RunE: func(cmd *cobra.Command, args []string) error {
		chunkSize, err := cmd.Flags().GetInt64(FlagChunkSize)
		if err != nil {
			return err
		}
		if !fsoper.FileExists(args[0]) {
			return fmt.Errorf("file is not exists %s , Please check this", args[0])
		}
		if !fsoper.FileExists(args[1]) {
			return fmt.Errorf("file is not exists %s , Please check this", args[0])
		}

		return RunERHA(args[0], args[1], chunkSize)
	},
}

func init() {
	rootCmd.AddCommand(
		GenSigCmd(),
		DeltaCmd(),
	)
	rootCmd.PersistentFlags().Int64(FlagChunkSize, types.DefaultChunkSize, "chunk_size for split the data into chunks, Default: 2")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
