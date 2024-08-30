-- Create ENUM type first
CREATE TYPE status_order AS ENUM (
    'pending',
    'delivering',
    'completed',
    'cancelled'
);

-- Create the orders table
CREATE TABLE orders (
                        id UUID DEFAULT uuid_generate_v4(),
                        cart_id UUID NOT NULL,
                        user_id UUID, -- Added user_id column
                        curier_id UUID, -- Note: "curier" seems like a typo for "courier", ensure this is intentional
                        status status_order NOT NULL DEFAULT 'pending',
                        created_at TIMESTAMP DEFAULT now(),
                        updated_at TIMESTAMP DEFAULT now(),
                        deleted_at TIMESTAMP DEFAULT NULL,

                        FOREIGN KEY (user_id) REFERENCES users(id) -- Make sure 'users' table exists
);
