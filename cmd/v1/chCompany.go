package v1

import (
	"fmt"

	"github.com/spf13/cobra"
)

// chCompanyCmd represents the chCompany command
var chCompanyCmd = &cobra.Command{
	Use:   "chCompany",
	Short: "Change identifier for a company",
	Long: `Change the siret identifier of on company`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("chCompany called")
	},
}

func init() {
	v1Cmd.AddCommand(chCompanyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// chCompanyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// chCompanyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
