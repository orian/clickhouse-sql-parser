
        CREATE TABLE fact_1000p_cached (p UInt16, id UInt64)
        ENGINE = MergeTree PARTITION BY p ORDER BY id
        SETTINGS refresh_statistics_interval = 1
    