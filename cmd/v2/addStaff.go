package v2

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"devoir10_ravet/model"
)

// addStaffCmd represents the insrow command
var addStaffCmd *cobra.Command

func init() {
	v2Cmd.AddCommand(addStaffCmd)

	addStaffCmd = &cobra.Command{
		Use:     "addStaff",
		Short:   "addStaff is used to add staff in database",
		Long:    "addStaff is used to add staff in database by specifying is <siret> and <insee>",
		PreRunE: addStaffPreRun,
		Run: func (cmd *cobra.Command, args []string) {
			fmt.Println("Running addStaff command")
		},

		// C'est à vous
	}
}

func addStaffPreRun(cmd *cobra.Command, args []string) error {
	if err := cobra.ExactArgs(2)(cmd, args); err != nil {
		return err
	}

	siret := args[0] // Two there should be. No more, no less.
	if siret == "" {
		return errSIRETIsEmpty
	}
	insee := args[1]
	if insee == "" {
		return errINSEEIsEmpty
	}

	ctx := context.Background()
	db := mustOpenDB()
	tx, err := db.BeginTx(ctx, model.TxOptions)
	if err != nil {
		log.Fatalf("creating staff insertion transaction: %v\n", err)
	}

	company, err := model.GetCompanyWithoutStaff(ctx, tx, siret)
	if err != nil {
		log.Fatal(err)
	}
	product, err := model.GetPerson(ctx, tx, insee)
	if err != nil {
		log.Fatal(err)
	}
	addStaffCmd.Run = func(_ *cobra.Command, _ []string) {
		var err error
		defer func() {
			db.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
		err = addStaff(ctx, tx, company, product)
	}

	return nil
}

func addStaff(ctx context.Context, tx *sql.Tx, company model.Company, person model.Person) error {
/*	// Attempt hiring.
	ins := `
	C'est à vous
   	(?, ?, ?);
	`
	res, err := tx.ExecContext(ctx, ins, append([]interface{}{company.SIRET}, person.SQLValues()...)...)
	if err != nil {

	//	C'est à vous

	// Report on insertion.
	if commitErr := tx.Commit(); commitErr != nil {
		err = fmt.Errorf("committing hiring transaction: %v", commitErr)
	}*/
	return err
}
