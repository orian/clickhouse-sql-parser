
        CREATE TABLE rmt_final_j (p UInt16, id UInt64, version UInt64, status String)
        ENGINE = ReplacingMergeTree(version) PARTITION BY p ORDER BY id
        SETTINGS refresh_statistics_interval = 0
    