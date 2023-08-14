CREATE TABLE students (
                              id serial not null unique,
                              name varchar(255) not null,
                              lastname varchar(255) not null,
                              isu integer not null unique,
                              groupNumber varchar(255) not null
);
CREATE TABLE groups (
    number varchar(255) not null,
    faculty varchar(255) not null,
    course  integer not null
);