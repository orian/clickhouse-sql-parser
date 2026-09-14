
        CREATE TABLE test_pk_long_back
        (
            `dt` DateTime,
            `key` String
        )
        ENGINE = MergeTree
        ORDER BY (dt, key || 'a', key || 'b', key || 'c', key || 'd', key || 'e', key || 'f', key || 'g', key || 'h', key || 'i', key || 'j', key || 'k', key || 'l', key || 'm', key || 'n', key || 'o', key || 'p', key || 'q', key || 'r', key || 's', key || 't', key)
        SETTINGS index_granularity = 4;
    