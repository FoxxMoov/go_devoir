package v1

import (
	"fmt"

	"github.com/spf13/cobra"
)

// lsCompanyCmd represents the lsCompany command
var lsCompanyCmd = &cobra.Command{
	Use:   "lsCompany",
	Short: "List companies and their staff",
	Long: `List all compagnies and all their staff`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("lsCompany called")
	},
}

func init() {
	v1Cmd.AddCommand(lsCompanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lsCompanyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lsCompanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
