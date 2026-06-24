CREATE SCHEMA myapp;

CREATE TABLE myapp.students(
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    fio VARCHAR(100) NOT NULL CHECK(char_length(fio) BETWEEN 3 AND 100),
    student_group VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);