CREATE TABLE student
(
    id        bigserial   not null primary key,
    name      varchar(24) not null,
    birthday  date        not null,
    school_id bigint      not null
);
