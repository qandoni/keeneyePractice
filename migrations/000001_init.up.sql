CREATE SCHEMA myapp;

CREATE TABLE myapp.users (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    login TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL CHECK (role IN ('student', 'teacher', 'admin'))
);

INSERT INTO myapp.users (
    login,
    password_hash,
    role
)
VALUES (
    'admin',
    '$2a$10$FeWE6q1VzFPiWZBRyj8kU.2bs8s3QY7M0QOaSbuBtumy9yryP137.',
    'admin'
);

CREATE TABLE myapp.refresh_tokens (
    id BIGSERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    user_id BIGINT NOT NULL,
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT refresh_tokens_user_id_fkey
        FOREIGN KEY (user_id)
        REFERENCES myapp.users(id)
        ON DELETE CASCADE,

    CONSTRAINT refresh_tokens_token_hash_key
        UNIQUE(token_hash),

    CONSTRAINT refresh_tokens_user_id_key
        UNIQUE(user_id)
);

CREATE TABLE myapp.groups (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE myapp.students (
    id SERIAL PRIMARY KEY,
    version INT NOT NULL DEFAULT 1,
    user_id INT NOT NULL UNIQUE REFERENCES myapp.users(id) ON DELETE CASCADE,
    group_id INT NOT NULL REFERENCES myapp.groups(id),
    fio VARCHAR(100) NOT NULL CHECK(char_length(fio) BETWEEN 3 AND 100),
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE myapp.teachers (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL UNIQUE REFERENCES myapp.users(id) ON DELETE CASCADE,
    version INT NOT NULL DEFAULT 1,
    fio VARCHAR(100) NOT NULL CHECK(char_length(fio) BETWEEN 3 AND 100),
    phone_number VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE myapp.teacher_groups (
    teacher_id INT NOT NULL REFERENCES myapp.teachers(id) ON DELETE CASCADE,
    version INT NOT NULL DEFAULT 1,
    group_id INT NOT NULL REFERENCES myapp.groups(id) ON DELETE CASCADE,

    PRIMARY KEY (teacher_id, group_id)
);

CREATE INDEX idx_students_group_id ON myapp.students(group_id);
CREATE INDEX idx_students_user_id ON myapp.students(user_id);

CREATE INDEX idx_teachers_user_id ON myapp.teachers(user_id);

CREATE INDEX idx_teacher_groups_teacher_id ON myapp.teacher_groups(teacher_id);
CREATE INDEX idx_teacher_groups_group_id ON myapp.teacher_groups(group_id);