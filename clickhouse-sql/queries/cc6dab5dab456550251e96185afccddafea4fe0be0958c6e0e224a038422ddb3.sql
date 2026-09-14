
        CREATE TABLE topk_join_l (id UInt64, k Int64, payload String) ENGINE = MergeTree() ORDER BY id SETTINGS auto_statistics_types=''
    