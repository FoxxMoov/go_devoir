-- +goose Up
SELECT 'up SQL query';
drop table staff;
create table people
(
    insee  integer(13) not null
        primary key,

    given   varchar(40),
    last    varchar(40)
);
create table staff
(
    insee integer(13) not null
        primary key
        constraint staff_people_insee_fk
            references people,
    siret integer not null
        references campanies
            on update restrict on delete restrict,
    name varchar(40),
    salary float
);
-- +goose Down
SELECT 'down SQL query';
drop table people;