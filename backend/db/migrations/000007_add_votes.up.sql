CREATE TABLE vote_types (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255) NOT NULL UNIQUE
);

INSERT INTO vote_types (type) VALUES ('up'), ('down'), ('invalid'), ('count');

CREATE TABLE votes (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    poll_option_id INTEGER REFERENCES poll_options(id) ON DELETE CASCADE,
    vote_type VARCHAR(255) NOT NULL REFERENCES vote_types(type),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT pk_votes PRIMARY KEY (user_id, poll_option_id)
);