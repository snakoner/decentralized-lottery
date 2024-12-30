CREATE TABLE IF NOT EXISTS chat (
    id                  serial          not null unique PRIMARY KEY,
    ts                  bigint          not null,
    user_addr           varchar(42)     not null,
    content             text            not null
);