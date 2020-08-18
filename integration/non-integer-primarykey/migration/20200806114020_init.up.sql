CREATE TABLE Student
(
    id        varchar(8)  not null primary key,
    name      varchar(24) not null,
    birthday  date        not null,
    school_id bigint      not null
);

-- some fake data
INSERT INTO Student(id, name, birthday, school_id)
VALUES ('20000101', 'A', '1990-01-01', 1),
       ('20000102', 'B', '1991-02-02', 2),
       ('20000103', 'C', '1992-03-03', 3),
       ('20000104', 'D', '1991-04-04', 1);