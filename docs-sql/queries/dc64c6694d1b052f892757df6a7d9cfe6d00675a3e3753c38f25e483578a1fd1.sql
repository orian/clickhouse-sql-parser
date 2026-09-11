-- Create a table
CREATE TABLE IF NOT EXISTS spark_catalog.db.time_travel_example_3 (
order_number int,
product_code string
  )
USING iceberg
OPTIONS ('format-version'='2');

ts = now();

-- Query the table at a specific timestamp
  SELECT * FROM spark_catalog.db.time_travel_example_3 TIMESTAMP AS OF ts; -- Finises with error: Cannot find a snapshot older than ts.
