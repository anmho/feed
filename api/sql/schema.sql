



CREATE TABLE users (
    user_id TEXT PRIMARY KEY,
    name TEXT NOT NULL
)
;

CREATE TABLE followers (
    from_user TEXT REFERENCES users (user_id),
    to_user TEXT REFERENCES users (user_id)
)
;


CREATE TABLE feeds (
    user_id TEXT NOT NULL,
    post_id TEXT NOT NULL
)
;

CREATE TABLE posts (
    id TEXT NOT NULL PRIMARY KEY,
    author_id TEXT NOT NULL REFERENCES users(user_id),
    content TEXT NOT NULL
)
;

