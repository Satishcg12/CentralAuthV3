-- +goose Up
-- +goose StatementBegin
-- Clients table (for OAuth applications)
CREATE TABLE clients (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    client_id VARCHAR(50) NOT NULL UNIQUE,
    client_secret VARCHAR(100) NOT NULL,
    redirect_uris TEXT[] NOT NULL,
    website_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    is_confidential BOOLEAN DEFAULT TRUE,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_clients_client_id ON clients(client_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
