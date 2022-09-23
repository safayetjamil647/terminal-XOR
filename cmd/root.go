package cmd

import (
    "context"
    // "errors"

    "github.com/spf13/cobra"
    // "github.com/spf13/viper"
    // "golang.org/x/sync/errgroup"
)

// Execute is the command line applications entry function
func Execute() error {
    rootCmd := &cobra.Command{
        Version: "v0.0.1",
        Use:     "xor",
		Long:    "xor for your devops environment",
        Example: "xor",
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
    }
    // Add sub commands
	rootCmd.AddCommand(initialize())
    
    return rootCmd.ExecuteContext(context.Background())
}

