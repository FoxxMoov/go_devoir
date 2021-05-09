-- +goose Up
SELECT 'up SQL query';
create table campanies
(
    siret        integer not null
        constraint campanies_pk
            primary key,
    date         datetime,
    total_salary float(3, 2)
);
create unique index campanies_siret_index
    on campanies (siret);
create table staff
(
    insee  integer(13) not null
        primary key,
    siret  integer     not null
        references campanies
            on update restrict on delete restrict,
    name   varchar(40),
    salary float
);
-- +goose Down
SELECT 'down SQL query';
drop table staff;
drop table campanies;
