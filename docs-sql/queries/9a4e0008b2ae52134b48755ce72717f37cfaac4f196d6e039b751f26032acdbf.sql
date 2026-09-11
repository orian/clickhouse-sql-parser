CREATE TABLE uk.uk_price_paid_simple_dist ON CLUSTER test_cluster
(
    date Date,
    town LowCardinality(String),
    street LowCardinality(String),
    price UInt32
)
ENGINE = Distributed('test_cluster', 'uk', 'uk_price_paid_simple', rand())
