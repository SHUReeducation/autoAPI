CREATE TABLE Student
(
    id        bigint auto_increment not null primary key,
    name      varchar(24)           not null,
    birthday  date                  not null,
    school_id bigint                not null
);

-- some fake data
INSERT INTO Student(name, birthday, school_id)
VALUES ('A', '1990-01-01', 1),
       ('B', '1991-02-02', 2),
       ('C', '1992-03-03', 3),
       ('D', '1991-04-04', 1);