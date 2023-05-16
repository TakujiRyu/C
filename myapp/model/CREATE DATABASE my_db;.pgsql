CREATE DATABASE my_db;

CREATE TABLE student (
    StdId int NOT NULL,
    FirstName varchar(45) NOT NULL,
    LastName varchar(45) DEFAULT NULL,
    Email varchar(45) NOT NULL,
    PRIMARY KEY (StdId),
    UNIQUE (Email)
)

SELECT my_db;

CREATE TABLE content (
    ContentId varchar(6) NOT NULL,
    Post varchar NOT NULL,
    PRIMARY KEY (ContentId)
)

DROP TABLE content