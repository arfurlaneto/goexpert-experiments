package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		repeat, _ := cmd.Flags().GetInt32("repeat")
		for i := 0; i < int(repeat); i++ {
			if name != "" {
				fmt.Printf("ping %s\n", name)
			} else {
				fmt.Printf("ping called\n")
			}
		}
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("called before the command")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("called after the command")
	},
}

var name string

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	pingCmd.Flags().Int32P("repeat", "r", 1, "How many times we should ping.")
	pingCmd.MarkFlagRequired("repeat")
	pingCmd.Flags().StringVarP(&name, "name", "n", "", "Name of who to ping.")
}
