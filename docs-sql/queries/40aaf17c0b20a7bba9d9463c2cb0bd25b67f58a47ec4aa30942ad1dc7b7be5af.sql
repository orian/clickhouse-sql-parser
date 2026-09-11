  -- sales テーブルを作成
  CREATE TABLE sales (
    date Date,
    product_id UInt32,
    product_name String,
    category String,
    quantity UInt32,
    revenue Decimal(10, 2),
    sales_amount Decimal(10, 2)
  )
  ENGINE = MergeTree()
  ORDER BY (date, product_id);

  -- サンプルデータを挿入
  INSERT INTO sales VALUES
  ('2024-01-05', 12345, 'Laptop Pro', 'Electronics', 2, 1799.98, 1799.98),
  ('2024-01-06', 12345, 'Laptop Pro', 'Electronics', 1, 899.99, 899.99),
  ('2024-01-10', 12346, 'Wireless Mouse', 'Electronics', 5, 124.95, 124.95),
  ('2024-01-15', 12347, 'USB-C Cable', 'Accessories', 10, 125.00, 125.00),
  ('2024-01-20', 12345, 'Laptop Pro', 'Electronics', 3, 2699.97, 2699.97),
  ('2024-01-25', 12348, 'Monitor 4K', 'Electronics', 2, 598.00, 598.00),
  ('2024-02-01', 12345, 'Laptop Pro', 'Electronics', 1, 899.99, 899.99),
  ('2024-02-05', 12349, 'Keyboard Mechanical', 'Accessories', 4, 319.96, 319.96),
  ('2024-02-10', 12346, 'Wireless Mouse', 'Electronics', 8, 199.92, 199.92),
  ('2024-02-15', 12350, 'Webcam HD', 'Electronics', 3, 164.85, 164.85);
