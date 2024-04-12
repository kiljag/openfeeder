
CREATE TABLE feed (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    feed_hash TEXT NOT NULL UNIQUE
);

create TABLE feed_item (
    id BIGSERIAL PRIMARY KEY,
    feed_id BIGINT NOT NULL REFERENCES feed(id),
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    description TEXT,
    pub_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    item_hash TEXT NOT NULL UNIQUE
);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE user_feed (
    user_id BIGINT NOT NULL REFERENCES users(id),
    feed_id BIGINT NOT NULL REFERENCES feed(id)
);

CREATE TABLE user_feed_item (
    user_id BIGINT REFERENCES users(id),
    feed_id BIGINT REFERENCES feed(id),
    item_id BIGINT REFERENCES feed_item(id),
    is_viewed BOOLEAN
);
