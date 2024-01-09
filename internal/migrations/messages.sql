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