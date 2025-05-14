-- +goose Up
-- +goose StatementBegin
-- User consents table
CREATE TABLE user_consents (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    scope TEXT[] NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, client_id)
);

CREATE INDEX idx_user_consents_user_id ON user_consents(user_id);
CREATE INDEX idx_user_consents_client_id ON user_consents(client_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_consents;
-- +goose StatementEnd
