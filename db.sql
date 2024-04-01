
CREATE TABLE feed (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT NOT NULL
);

create TABLE feed_item (
    id BIGSERIAL PRIMARY KEY,
    feed_id BIGINT NOT NULL REFERENCES feed(id),
    title TEXT NOT NULL,
    link TEXT NOT NULL,
    description TEXT,
    pubDate TIMESTAMP
);