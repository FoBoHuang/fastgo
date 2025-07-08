package app

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFastGOCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fg-apiserver",
		Short: "A very lightweight full go project",
		Long: `A very lightweight full go project, designed to help beginners quickly
        learn Go project development.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Hello FastGO!")
			return nil
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./fg-apiserver param1 param2
		Args: cobra.NoArgs,
	}

	return cmd
}
