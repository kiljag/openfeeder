
CREATE TABLE feed (
    id BIGSERIAL PRIMARY KEY,
    title TEXT,
    url TEXT,
    description TEXT
);

create TABLE feed_item (
    id BIGSERIAL PRIMARY KEY,
    feed_id BIGINT REFERENCES feed(id),
    title TEXT,
    description TEXT,
    link TEXT,
    pubDate TIMESTAMP
);