package cmd

import "github.com/spf13/cobra"

var root = &cobra.Command{
	Use: "gcsv",
}

func init() {
	root.AddCommand(generate)
}

func Execute() error {
	return root.Execute()
}
