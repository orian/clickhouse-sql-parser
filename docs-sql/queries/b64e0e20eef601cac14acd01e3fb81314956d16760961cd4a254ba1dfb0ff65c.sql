  -- 1. Crear una tabla para probar cadenas y números
  CREATE TABLE IF NOT EXISTS users (
      name String,
      age UInt8,
      salary Float64
  ) ENGINE = Memory;

  INSERT INTO users VALUES
      ('John Doe', 25, 75000.50),
      ('Jane Smith', 30, 85000.75),
      ('Peter Jones', 20, 50000.00);

  -- 2. Crear una tabla para probar fechas y marcas temporales
  CREATE TABLE IF NOT EXISTS events (
      event_date Date,
      event_timestamp DateTime
  ) ENGINE = Memory;

  INSERT INTO events VALUES
      ('2024-01-15', '2024-01-15 14:30:00'),
      ('2024-01-15', '2024-01-15 15:00:00'),
      ('2024-01-16', '2024-01-16 10:00:00');

  -- 3. Crear una tabla para probar arrays
  CREATE TABLE IF NOT EXISTS products (
      id UInt32,
      name String
  ) ENGINE = Memory;

  INSERT INTO products VALUES (1, 'Laptop'), (2, 'Monitor'), (3, 'Mouse'), (4, 'Keyboard');

  -- 4. Crear una tabla para probar Map (similar a struct)
  CREATE TABLE IF NOT EXISTS accounts (
      user_id UInt32,
      status String,
      type String
  ) ENGINE = Memory;

  INSERT INTO accounts VALUES
      (101, 'active', 'premium'),
      (102, 'inactive', 'basic'),
      (103, 'active', 'basic');

  -- 5. Crear una tabla para probar identificadores
  CREATE TABLE IF NOT EXISTS sales_2024 (
      value UInt32
  ) ENGINE = Memory;

  INSERT INTO sales_2024 VALUES (100), (200), (300);
