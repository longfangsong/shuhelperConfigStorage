create type ThemeMode as enum ('dark', 'light');

create type ServerOrClient as enum ('server', 'client');

create table Config
(
    student_id varchar(16) not null,
    mode       ThemeMode,
    saveTodoIn ServerOrClient
);

create unique index Config_student_id_uindex
    on Config (student_id);

alter table Config
    add constraint Config_pk
        primary key (student_id);

