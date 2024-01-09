CREATE TABLE users
(
    uid        UUID,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null,
    chat_id    BIGINT                                             not null,
    name       varchar(250)             default null,
    username   varchar(250)             default null
);

CREATE unique index bot_chat_id_idx on users (chat_id);

ALTER TABLE users
    ADD PRIMARY KEY (uid);