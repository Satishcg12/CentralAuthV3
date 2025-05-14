-- +goose Up
-- +goose StatementBegin
-- API scopes table
CREATE TABLE api_scopes (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_default BOOLEAN DEFAULT FALSE,
    requires_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default scopes
INSERT INTO api_scopes (name, display_name, description, is_default) VALUES 
('profile', 'Profile', 'Access to your basic profile information', true),
('email', 'Email', 'Access to your email address', true),
('openid', 'OpenID', 'Standard OpenID Connect scope', true),
('offline_access', 'Offline Access', 'Allows the application to access resources on your behalf even when you are offline', false);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS api_scopes;
-- +goose StatementEnd
