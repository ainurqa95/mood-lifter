-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE users
(
    id         serial,
    uid        UUID,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    chat_id    BIGINT                                             not null,
    name       varchar(250)             default null,
    username   varchar(250)             default null
);

CREATE unique index bot_chat_id_idx on users (chat_id);

ALTER TABLE users
    ADD PRIMARY KEY (id);

CREATE TABLE messages
(
    id         serial,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    text       text                                               not null,
    chat_id    BIGINT
);

CREATE index messages_chat_idx on messages (chat_id);

ALTER TABLE messages
    ADD PRIMARY KEY (id);

CREATE TABLE compliments
(
    text       text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE users;
DROP TABLE messages;
DROP TABLE compliments;