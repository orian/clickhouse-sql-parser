
        CREATE TABLE with_lonely
        (
        `id` UInt64,
        `dt` Date,
        `val` UInt64,
        `version` UInt64
        )
        ENGINE = ReplacingMergeTree(version)
        PARTITION BY dt
        ORDER BY (id);
    