create table products (
                                        id uuid primary key default uuid_generate_v4(),
    name varchar(255),
    description varchar(255),
    price numeric,
    image varchar(255),

    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp default null
    );

create index  idx_products_deleted_at on products (deleted_at);