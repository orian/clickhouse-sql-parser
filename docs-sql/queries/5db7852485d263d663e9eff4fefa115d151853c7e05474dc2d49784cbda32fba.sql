CREATE TABLE `users`
(
    `id`               Int32,
    `name`             Nullable(String),
    `status`           Nullable(String),
    `_fivetran_synced` DateTime64(9, 'UTC'),
    `_fivetran_start`  DateTime64(9, 'UTC'),
    `_fivetran_end`    Nullable(DateTime64(9, 'UTC')),
    `_fivetran_active` Nullable(Bool)
) ENGINE = SharedReplacingMergeTree('/clickhouse/tables/{uuid}/{shard}', '{replica}', _fivetran_synced)
ORDER BY (id, _fivetran_start)
SETTINGS index_granularity = 8192
