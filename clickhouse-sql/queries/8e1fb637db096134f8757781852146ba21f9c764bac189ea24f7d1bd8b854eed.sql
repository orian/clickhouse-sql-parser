
        CREATE TABLE test_pk
        (
            `dt` DateTime,
            `key` String
        )
        ENGINE = MergeTree
        ORDER BY (dt, key)
        SETTINGS index_granularity = 4;
    