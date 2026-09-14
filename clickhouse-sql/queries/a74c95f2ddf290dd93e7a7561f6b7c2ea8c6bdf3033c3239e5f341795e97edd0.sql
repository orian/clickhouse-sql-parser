
        CREATE TABLE test_pk_long
        (
            `dt` DateTime,
            `key` String
        )
        ENGINE = MergeTree
        ORDER BY (dt, key, key || 'a', key || 'b', key || 'c', key || 'd', key || 'e', key || 'f', key || 'g', key || 'h', key || 'i', key || 'j', key || 'k', key || 'l', key || 'm', key || 'n', key || 'o', key || 'p', key || 'q', key || 'r', key || 's', key || 't')
        SETTINGS index_granularity = 4;
    