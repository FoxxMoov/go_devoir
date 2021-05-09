package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upMigration, downMigration)
}

const createStaff = `create table staff
(
    insee  integer(13) not null
        primary key,
    siret  integer     not null
        references campanies
            on update restrict on delete restrict,
    name   varchar(40),
    salary float
);`

const createCompanies = `create table campanies
(
    siret        integer not null
        constraint campanies_pk
            primary key,
    date         datetime,
    total_salary float(3, 2)
);
create unique index campanies_siret_index
    on campanies (siret);`

const dropStaff = `drop table staff;`
const dropCompanies = `drop table campanies;`

func upMigration(tx *sql.Tx) error {
	return SimpleRun(tx, []string {createStaff, createCompanies})
}

func downMigration(tx *sql.Tx) error {
	return SimpleRun(tx, []string {dropStaff, dropCompanies})
}
