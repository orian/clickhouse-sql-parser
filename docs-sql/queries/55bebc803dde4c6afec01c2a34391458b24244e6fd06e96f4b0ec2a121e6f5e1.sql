CREATE TABLE uk.uk_price_paid_simple_partitioned
(
  date Date,
  town LowCardinality(String),
  street LowCardinality(String),
  price UInt32
)
ENGINE = MergeTree
ORDER BY (town, street)
PARTITION BY toStartOfMonth(date)
