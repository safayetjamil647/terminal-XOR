package cmd

import (
    // tea "github.com/charmbracelet/bubbletea"
    // "github.com/safayetjamil647/terminal-xor/cmd/internal/tui"
    "github.com/spf13/cobra"
    // "github.com/spf13/viper"
)

func initialize() *cobra.Command {
    init := &cobra.Command{
        Use:     "initialize",
        Short:   "init the xor config.",
        Long:    "init provision the xor configuration.",
        Example: "rkl init",
        Aliases: []string{"i", "init"},
        RunE: func(cmd *cobra.Command, args []string) error {
            return nil
        },
    }
    return init
}