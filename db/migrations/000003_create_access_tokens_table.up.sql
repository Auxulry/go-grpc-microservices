CREATE TABLE IF NOT EXISTS access_tokens (
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY NOT NULL,
    user_id uuid NOT NULL,
    revoked BOOLEAN DEFAULT false,
    expires_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
);