CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255),
    song_name VARCHAR(255),
    release VARCHAR(255),
    text TEXT,
    link VARCHAR(255)
);