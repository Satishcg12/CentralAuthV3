-- +goose Up
-- +goose StatementBegin
-- Access tokens table
CREATE TABLE access_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token VARCHAR(255) NOT NULL UNIQUE,
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    scope TEXT[] NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_access_tokens_token ON access_tokens(token);
CREATE INDEX idx_access_tokens_user_id ON access_tokens(user_id);
CREATE INDEX idx_access_tokens_client_id ON access_tokens(client_id);
CREATE INDEX idx_access_tokens_expires_at ON access_tokens(expires_at);

-- Refresh tokens table
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token VARCHAR(255) NOT NULL UNIQUE,
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    scope TEXT[] NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_refresh_tokens_token ON refresh_tokens(token);
CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_client_id ON refresh_tokens(client_id);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS access_tokens;
-- +goose StatementEnd
