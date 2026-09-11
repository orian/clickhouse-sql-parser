
        CREATE TABLE topk_join_r (id UInt64, value String) ENGINE = MergeTree() ORDER BY id SETTINGS auto_statistics_types=''
    