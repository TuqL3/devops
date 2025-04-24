CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(50),
    active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tạo một số dữ liệu mẫu
INSERT INTO users (name, email, password, role, active)
VALUES 
    ('Admin User', 'admin@example.com', '$2a$10$hKDVYxLefVHV/vV76Algb.tOyCkz1C1UU5Myz0Qx2qgS4ZzIHzZSO', 'admin', true),
    ('Test User', 'user@example.com', '$2a$10$hKDVYxLefVHV/vV76Algb.tOyCkz1C1UU5Myz0Qx2qgS4ZzIHzZSO', 'user', true);