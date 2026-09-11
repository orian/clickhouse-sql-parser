CREATE TABLE product_sales
(
    product_name String,
    price Decimal32(2),
    sales_count UInt32
) ENGINE = Memory;

INSERT INTO product_sales VALUES
    ('Laptop', 999.99, 10),
    ('Phone', 499.99, 15),
    ('Tablet', 299.99, 0),
    ('Watch', 1199.99, 5),
    ('Headphones', 79.99, 20);

SELECT argMaxIf(product_name, price, sales_count >= 10) AS most_expensive_popular_product
FROM product_sales;
