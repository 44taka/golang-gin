CREATE TABLE IF NOT EXISTS tests(
    id SERIAL NOT NULL PRIMARY KEY,
    name varchar(100) NOT NULL,
    password varchar(10) NOT NULL
);