package v2

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chPersonCmd represents the chPerson command
var chPersonCmd = &cobra.Command{
	Use:   "chPerson",
	Short: "Update the name of a person",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chPerson called")
	},
}

func init() {
	v2Cmd.AddCommand(chPersonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chPersonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chPersonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
