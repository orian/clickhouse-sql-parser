CREATE TABLE db.table_local on cluster cluster (
    `id` UInt64,
    `created_at` DateTime,
    `item` String
)
    ENGINE = ReplacingMergeTree
    ORDER BY (id, created_at);

CREATE TABLE db.table on cluster cluster (
    `id` UInt64,
    `created_at` DateTime,
    `item` String
)
    ENGINE = Distributed ('cluster', 'db', 'table_local', cityHash64(id));
