-- +goose Up
-- +goose StatementBegin
-- Authorization codes table
CREATE TABLE authorization_codes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code VARCHAR(100) NOT NULL UNIQUE,
    client_id UUID NOT NULL REFERENCES clients(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    redirect_uri TEXT NOT NULL,
    scope TEXT[] NOT NULL,
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_authorization_codes_code ON authorization_codes(code);
CREATE INDEX idx_authorization_codes_user_id ON authorization_codes(user_id);
CREATE INDEX idx_authorization_codes_client_id ON authorization_codes(client_id);
CREATE INDEX idx_authorization_codes_expires_at ON authorization_codes(expires_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS authorization_codes;
-- +goose StatementEnd
