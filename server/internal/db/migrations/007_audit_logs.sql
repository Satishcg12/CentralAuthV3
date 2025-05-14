-- +goose Up
-- +goose StatementBegin
-- Audit logs table
CREATE TYPE audit_action AS ENUM (
    'login', 
    'logout', 
    'register', 
    'password_reset', 
    'email_verification', 
    'token_issue', 
    'token_revoke',
    'client_create',
    'client_update',
    'client_delete',
    'consent_grant',
    'consent_revoke'
);

CREATE TABLE audit_logs (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    client_id UUID REFERENCES clients(id) ON DELETE SET NULL,
    action audit_action NOT NULL,
    ip_address VARCHAR(45),
    user_agent TEXT,
    details JSONB,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_client_id ON audit_logs(client_id);
CREATE INDEX idx_audit_logs_action ON audit_logs(action);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS audit_logs;
DROP TYPE IF EXISTS audit_action;
-- +goose StatementEnd
