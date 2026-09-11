CREATE TABLE default.orders
(
    `id` Int64,
    `region` String,
    `amount` Float64
)
ENGINE = MergeTree()
ORDER BY (region, id);
