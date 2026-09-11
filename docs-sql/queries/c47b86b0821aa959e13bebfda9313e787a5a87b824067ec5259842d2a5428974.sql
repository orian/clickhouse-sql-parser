  -- Crear la tabla products
  CREATE TABLE products (
      product_id UInt32,
      product_name String,
      price Decimal(10, 2)
  )
  ENGINE = MergeTree()
  ORDER BY product_id;

  -- Insertar datos de muestra
  INSERT INTO products (product_id, product_name, price) VALUES
  (1, 'Laptop', 899.99),
  (2, 'Wireless Mouse', 24.99),
  (3, 'USB-C Cable', 12.50),
  (4, 'Monitor', 299.00),
  (5, 'Keyboard', 79.99),
  (6, 'Webcam', 54.95),
  (7, 'Desk Lamp', 34.99),
  (8, 'External Hard Drive', 119.99),
  (9, 'Headphones', 149.00),
  (10, 'Phone Stand', 15.99);
