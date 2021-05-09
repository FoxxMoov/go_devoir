/*
package migrations contains the numbered migrations used by Goose, and the
SimpleRun() utility function which Go migrations can use to simplify their logic.
 */
package migrations

import (
	"database/sql"

	"devoir10_ravet/model"
)

type strings []string

/*
People is used during the v1 to v2 migration to load persons.
*/
var People model.People

/*
SimpleRun is a trivial wrapper executing SQL instructions without a result in sequence.
*/
func SimpleRun(tx *sql.Tx, instructions strings) error {
	for _, instruction := range instructions {
		_, err := tx.Exec(instruction)
		if err != nil {
			return err
		}
	}

	return nil
}
