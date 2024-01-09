create table compliments
(
    text       text,
    created_at timestamp with time zone default CURRENT_TIMESTAMP not null
);