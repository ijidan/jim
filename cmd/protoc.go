package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// protocCmd represents the protoc command
var protocCmd = &cobra.Command{
	Use:   "protoc",
	Short: "build command for protoc buffers",
	Long:  `build command for protoc buffers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("protoc called")
	},
}

func init() {
	rootCmd.AddCommand(protocCmd)
}
