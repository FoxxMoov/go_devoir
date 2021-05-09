package v2

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmPersonCmd represents the rmPerson command
var rmPersonCmd = &cobra.Command{
	Use:   "rmPerson",
	Short: "Delete an employee",
	Long: `Delete an employee and remove it from all compagnies`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmPerson called")
	},
}

func init() {
	v2Cmd.AddCommand(rmPersonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmPersonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmPersonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
