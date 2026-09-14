CREATE TABLE default.`clickhouse-local-ab404c86-56cc-495b-ad1d-fb343cac3bc0events`
(
    `favoriteDatabase` String DEFAULT 'ClickHouse',
    `date` Date DEFAULT today(),
    `name` String,
    `favoriteColor` String DEFAULT 'Yellow',
    `favoriteNumber` Float64 DEFAULT 21
)
ENGINE = MergeTree
ORDER BY date
