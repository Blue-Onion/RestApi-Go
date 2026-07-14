-- +goose Up
CREATE TABLE sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    refresh_token_hash TEXT NOT NULL,

    user_agent TEXT,
    ip_address INET,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_used_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,

    revoked_at TIMESTAMPTZ
);

CREATE INDEX idx_sessions_user_id
ON sessions(user_id);

CREATE INDEX idx_sessions_refresh_hash
ON sessions(refresh_token_hash);

-- +goose Down
DROP TABLE sessions;
