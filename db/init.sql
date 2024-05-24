CREATE TABLE posts (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       comments_enabled BOOLEAN DEFAULT TRUE,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE comments (
                          id SERIAL PRIMARY KEY,
                          post_id INTEGER NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
                          parent_id INTEGER REFERENCES comments(id) ON DELETE CASCADE,
                          content TEXT NOT NULL CHECK (LENGTH(content) <= 2000),
                          created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
