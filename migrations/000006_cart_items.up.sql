CREATE TABLE cart_products (
                               cart_id UUID NOT NULL,
                               product_id UUID NOT NULL,
                               quantity INT,

                               FOREIGN KEY (cart_id) REFERENCES carts(id),
                               FOREIGN KEY (product_id) REFERENCES products(id)
);
