-- Create ENUM type first
CREATE TYPE status AS ENUM (
    'PENDING',
    'ORDERED'
);

-- Create table with corrected syntax
CREATE TABLE carts (
                       id UUID DEFAULT uuid_generate_v4(),
                       user_id UUID NOT NULL,
                       status status NOT NULL DEFAULT 'PENDING',
                       created_at TIMESTAMP DEFAULT now(),
                       updated_at TIMESTAMP DEFAULT now(),
                       deleted_at TIMESTAMP DEFAULT NULL,
                       PRIMARY KEY (id)
);
