package v1

import (
	"fmt"

	"github.com/spf13/cobra"
)

// v1Cmd represents the v1 command
var v1Cmd = &cobra.Command{
	Use:   "v1",
	Short: "Load migration v1",
	Long: `Load migration v1 : 
	- v1 load
	- v1 lsCompany
	- v1 chCompany <siret1> <siret2>
	- v1 rmCompany
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1 called")
	},
}

func init() {
	v1Cmd.AddCommand(v1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// v1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// v1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
