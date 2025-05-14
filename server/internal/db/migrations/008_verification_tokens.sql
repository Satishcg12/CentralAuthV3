-- +goose Up
-- +goose StatementBegin
-- Verification tokens table for email verification, password reset, etc.
CREATE TABLE verification_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    token VARCHAR(100) NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type VARCHAR(20) NOT NULL, -- 'email_verification', 'password_reset', etc.
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    used BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_verification_tokens_token ON verification_tokens(token);
CREATE INDEX idx_verification_tokens_user_id ON verification_tokens(user_id);
CREATE INDEX idx_verification_tokens_expires_at ON verification_tokens(expires_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS verification_tokens;
-- +goose StatementEnd
