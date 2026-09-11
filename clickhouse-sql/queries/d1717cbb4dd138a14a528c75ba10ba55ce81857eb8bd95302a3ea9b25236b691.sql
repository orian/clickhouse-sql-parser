
        CREATE TABLE mv_dst (`key` Int64, `value` String)
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS non_replicated_deduplication_window=1000;
    