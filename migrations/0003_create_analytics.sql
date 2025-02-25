CREATE TABLE analytics (
    warehouse_id UUID NOT NULL,
    product_id UUID NOT NULL,
    sold_quantity INT NOT NULL,
    total_sum DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY (warehouse_id, product_id)
);
