CREATE TABLE clickhouse.default.my_table (
  id INT,
  name STRING
) USING ClickHouse
TBLPROPERTIES (
  engine = 'MergeTree()',
  order_by = 'id',
  'settings.allow_nullable_key' = '1',
  'settings.index_granularity' = '8192'
)
