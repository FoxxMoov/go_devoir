package v2

import (
	"fmt"

	"github.com/spf13/cobra"
)

// v2Cmd represents the v2 command
var v2Cmd = &cobra.Command{
	Use:   "v2",
	Short: "Load migration v2",
	Long: `Load migration v2 : 
	- v2 addStaff <siret> <insee>
	- v2 rmPerson <insee>
	- v2 chPerson <insee> "<given> <last>"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v2 called")
	},
}

func init() {
	v2Cmd.AddCommand(v2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// v2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// v2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
