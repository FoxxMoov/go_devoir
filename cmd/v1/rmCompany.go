package v1

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCompanyCmd represents the rmCompany command
var rmCompanyCmd = &cobra.Command{
	Use:   "rmCompany",
	Short: "Delete company",
	Long: `Delete one company and all their staff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmCompany called")
	},
}

func init() {
	v1Cmd.AddCommand(rmCompanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCompanyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCompanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
